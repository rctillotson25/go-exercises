package sorting

import (
	"algorithms/algoutils"
)

/*
	---------------------------
	Insertion Sorting Algorithm
	---------------------------

	* Space complexity: O(n) total - O(1) auxiliary. In place sorting, so no additional space is necessary.
	* Best-case time: O(n) - occurs when the array is already sorted.
	* Average-case time: O(n^2)
	* Worst-case time: O(n^2)
*/
func InsertionSort(A []int) {
	for i := 1; i < len(A); i++ {
		for j := i; j > 0; j--  {
			if A[j] < A[j-1] {
				algoutils.Swap(A, j, j-1)
			}
		}
	}
}


