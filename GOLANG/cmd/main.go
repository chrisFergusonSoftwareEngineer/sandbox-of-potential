package main

import (
	"fmt"

	"github.com/chrisFergusonSoftwareEngineer/sandbox-of-potential/GOLANG/locking"
)

func main() {
	fmt.Println("Running the SANDBOX OF POTENTIAL!!!")

	locking.TryOnceLocking()

	//using defer to (hopefully) ensure other threads complete first.
	defer fmt.Println("Exiting the SANDBOX OF POTENTIAL!!!")
}
