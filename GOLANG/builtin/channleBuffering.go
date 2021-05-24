package builtin

import "fmt"

func ChannelingTest() {
	// unbuffered := make(chan string)

	// unbuffered <- "Unbuffered "
	// unbuffered <- "test."

	// fmt.Printf(<-unbuffered) // causes PANIC
	// fmt.Printf(<-unbuffered)

	// fmt.Println()
	// fmt.Println()

	buffered := make(chan string, 2)

	buffered <- "Buffered "
	buffered <- "Test."
	//buffered <- "overflow"  // causes PANIC

	fmt.Printf(<-buffered)
	fmt.Printf(<-buffered)

	fmt.Println()
}
