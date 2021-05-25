package threadpool

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type SearchTask struct {
	Id         string
	Date       string
	ResultChan chan int
}

func (search *SearchTask) Do() {
	time.Sleep(2 * time.Second)
	result := rand.Int() % 10
	if result%2 == 1 {
		fmt.Println("ERROR: result odd [" + strconv.Itoa(result) + "]")
		close(search.ResultChan)
	} else {
		search.ResultChan <- result
	}
}

type WriteTask struct {
	Id     string
	Date   string
	result int
}

func (write *WriteTask) Do() {
	var recheckTask SearchTask
	recheckTask.Id = write.Id
	recheckTask.Date = write.Date
	recheckTask.ResultChan = make(chan int)

	//dispatch recheck task
	GetWorkerPool().Enqueue(&recheckTask)

	checkResult, resultFound := <-recheckTask.ResultChan
	if resultFound {
		fmt.Println("WRITE ERROR: Found result [" + strconv.Itoa(checkResult) + "]")
	}

	// write result to data store.
}
