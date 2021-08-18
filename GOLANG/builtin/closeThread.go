package builtin

import (
	"fmt"
	"strconv"
)

func CloseThreadExpirment() {
	testChan := make(chan int, 1)

	go func(neglectChanel chan int) {
		defer close(neglectChanel)
		//neglectChanel <- 1
		return
	}(testChan)

	result, status := <-testChan

	if status {
		fmt.Println("Result: " + strconv.Itoa(result))
	} else {
		fmt.Println("### NO RESULT")
	}
}
