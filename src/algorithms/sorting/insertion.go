package sorting

import (
	"fmt"
	"algorithms/algoutils"
)

/*
	---------------------------
	Insertion Sorting Algorithm
	---------------------------

	* Space complexity:
	* Best-case time: 
	* Average-case time: 
	* Worst-case time: N^2
*/
func InsertionSort(A []int) {
	fmt.Println("Insertion sort..")
	for i := 1; i < len(A); i++ {
		for j := i; j > 0; j--  {
			if A[j] < A[j-1] {
				algoutils.Swap(A, j, j-1)
			}
		}
	}
	fmt.Println(A)
}


