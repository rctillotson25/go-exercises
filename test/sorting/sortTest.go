package main

import (
	"algorithms/sorting"
	"random"
	"fmt"
)

const MAX_VALUE  = 100 
const NUM_VALUES = 50
func main() {

	A := random.RandIntN(NUM_VALUES, MAX_VALUE)
	fmt.Println(A)
	sorting.MergeSort()
	insertA := A
	sorting.InsertionSort(insertA)
	fmt.Println(A)
}
