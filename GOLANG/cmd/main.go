package main

import (
	"fmt"

	"github.com/chrisFergusonSoftwareEngineer/sandbox-of-potential/GOLANG/temporal"
)

func main() {
	fmt.Println("Running the SANDBOX OF POTENTIAL!!!")

	temporal.TryBigSeconds()

	//using defer to (hopefully) ensure other threads complete first.
	defer fmt.Println("Exiting the SANDBOX OF POTENTIAL!!!")
}
