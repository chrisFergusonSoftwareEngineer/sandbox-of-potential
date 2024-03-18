package builtin

import "fmt"

// a non-python file.
func TestArrayPointerFunction() {
	var valueTesting *[]int

	for n := 0; n < 3; n++ {
		valueTesting = updateValueArray(valueTesting)
	}

	for i := range *valueTesting {
		fmt.Printf("%d :: %d\n", i, (*valueTesting)[i])
	}
}

func updateValueArray(inputArray *[]int) *[]int {
	if inputArray == nil {
		inputArray = new([]int)
		*inputArray = make([]int, 3)
	}
	for i := range *inputArray {
		(*inputArray)[i] = (*inputArray)[i] + i
	}

	return inputArray
}
