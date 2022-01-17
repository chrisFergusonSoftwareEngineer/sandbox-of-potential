package main

import (
	"fmt"

	"github.com/chrisFergusonSoftwareEngineer/sandbox-of-potential/GOLANG/builtin"
)

func main() {
	fmt.Println("Running the SANDBOX OF POTENTIAL!!!")

	//using defer to (hopefully) ensure other threads complete first.
	defer fmt.Println("Exiting the SANDBOX OF POTENTIAL!!!")

	builtin.TestArrayPointerFunction()
}

// Testing git changes.
