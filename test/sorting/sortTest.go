package main

import (
	"algorithms/sorting"
	"fmt"
	"random"
	"os"
)

const MAX_VALUE = 100
const NUM_VALUES = 5000

func main() {
	mergeSortTest()
	insertionSortTest()
}

func mergeSortTest() {
	a := random.RandIntN(NUM_VALUES, MAX_VALUE)
	fmt.Println(a)
	fmt.Println("Running merge sort...")
	sorting.MergeSort(a)
	fmt.Println(a)
}

func insertionSortTest() {
	a := random.RandIntN(NUM_VALUES, MAX_VALUE)
	fmt.Println(a)
	fmt.Println("Running insertion sort...")
	sorting.InsertionSort(a)
	fmt.Println(a)
}
