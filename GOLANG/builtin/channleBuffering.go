package builtin

import (
	"fmt"
	"strconv"
)

func ChannelingTest() {
	// unbuffered := make(chan string)

	// unbuffered <- "Unbuffered " // causes PANIC
	// unbuffered <- "test."

	// fmt.Printf(<-unbuffered)
	// fmt.Printf(<-unbuffered)

	// fmt.Println()
	// fmt.Println()

	buffered := make(chan string, 2)

	fmt.Println(strconv.Itoa(len(buffered)) + " of " + strconv.Itoa(cap(buffered)))

	buffered <- "Buffered "
	fmt.Println(strconv.Itoa(len(buffered)) + " of " + strconv.Itoa(cap(buffered)))
	buffered <- "Test."
	fmt.Println(strconv.Itoa(len(buffered)) + " of " + strconv.Itoa(cap(buffered)))
	// buffered <- "overflow" // causes PANIC

	fmt.Printf(<-buffered)
	fmt.Printf(<-buffered)

	fmt.Println()
}
