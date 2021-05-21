package main

import (
	"fmt"

	"github.com/chrisFergusonSoftwareEngineer/sandbox-of-potential/GOLANG/json"
)

func main() {
	fmt.Println("Running the SANDBOX OF POTENTIAL!!!")

	json.ParseJsonWithGabs()

	//using defer to (hopefully) ensure other threads complete first.
	defer fmt.Println("Exiting the SANDBOX OF POTENTIAL!!!")
}

// Testing git changes.
