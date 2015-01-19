/*
   Tests for the sorting class.
   
   Note that it would probably be better in practice
   to keep all of these algorithm tests in separate functions,
   but this method allows the test to be expanded easily. Plus,
   being relatively new to Go, I wanted practice with passing 
   functions as parameters.

   If more data sets are needed, just create a type for holding 
   multiple sets and and rewrite part of the runTest algorithm to 
   runTests over all of the data sets instead of just arrOrig.
*/

package sorting

import (
	"testing"
	"utils"
)

// Standard int sorting case
var arrOrig = []int{9, 7, 2, 1, 4, 5, 0}
var arrSorted = []int{0, 1, 2, 4, 5, 7, 9}

type sortAlg struct {
	name string
	alg  func([]int) []int
}

type sortData struct {
	arrOrig []int
	arrSorted []int
}

func TestSortingAlgorithms(t *testing.T) {

	// algorithms to be tested
	algs := []sortAlg{
		{"Insertion Sort", InsertionSort},
		{"Merge Sort", MergeSort},
	}

	for i := range algs {
		runTest(algs[i], t)
	}
}



func runTest(a sortAlg, t *testing.T) {

	// test cases for algorithms
	data := []sortData{
		{[]int{9,7,2,1,4,5,0}, []int{0,1,2,4,5,7,9}},
		{[]int{}, []int{}},
	}

	for i := range data {
		// initialize new slice for sorting
		var arr = make([]int, len(data[i].arrOrig))
		copy(arr, data[i].arrOrig)

		arr = a.alg(arr)

		if !utils.Equal(arr, data[i].arrSorted) {
			t.Error("\nAlgorithm: ", a.name, "\nExpected: ", data[i].arrSorted, "\nresult: ", arr)
		}
	}
}
