package sorting

import (
	"testing"
)

// Standard int sorting case 
var arrOrig = []int{9,7,2,1,4,5,0}
var arrSorted = []int{0,1,2,4,5,7,9}

// Edge Cases
// TODO - need to add edge cases later.

func TestInsertionSort(t *testing.T) {

	// must copy over the array so arrOrig isn't modified. 
	var arr = make([]int, len(arrOrig))
	copy(arr, arrOrig)

	arr = InsertionSort(arr)

	if !compareArrays(arr, arrSorted) {
		t.Error("Expected: ", arrSorted, " result: ", arr)
	}

}


func TestMergeSort(t *testing.T) {

	var arr = make([]int, len(arrOrig))
	copy(arr, arrOrig)

	arr = MergeSort(arr)

	if !compareArrays(arr, arrSorted) {
		t.Error("Expected: ", arrSorted, " result: ", arr)
	}
}

/*
   Return true if two arrays (or slices)  of integers
   are identical (same order, length, values).
*/
func compareArrays(a []int, b []int) bool {

	if (len(a) != len(b)) {
		return false
	} else {
		for i := range(a) {
			if a[i] != b[i] {
				return false
			}
		}
	}
	return true

}


