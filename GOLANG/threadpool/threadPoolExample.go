package threadpool

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func ThreadPoolDemo() {
	GetWorkerPool().Start()
	var exitGroup sync.WaitGroup

	for i := 0; i < 10000; i++ {
		id := "Test--" + strconv.Itoa(i)
		offset := i % 7
		date := time.Now().Add(time.Hour * 24 * -1 * time.Duration(offset)).Format("2006-01-02")

		query := QueryForRollup{
			Id:     id,
			Date:   date,
			Result: make(chan int),
		}

		exitGroup.Add(1)
		GetWorkerPool().Enqueue(&query)

		go func(task QueryForRollup, groupComplete *sync.WaitGroup) {
			value, found := <-task.Result
			if found {
				fmt.Println("FOUND [" + strconv.Itoa(value) + "] for [" + task.Id + ", " + task.Date + "]")
			} else {
				fmt.Println("Channel closed.")
			}
			groupComplete.Done()
		}(query, &exitGroup)
	}

	exitGroup.Wait()
}
