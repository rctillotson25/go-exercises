package sorting

import (
	"algorithms/algoutils"
)

/*
	---------------------------
	Merge Sorting Algorithm
	---------------------------

	* Space complexity: O(n) total - O(n) auxiliary.
	* Best-case time: O(n) - occurs when the array is already sorted.
	* Average-case time: O(n^2)
	* Worst-case time: O(n^2)
*/
func MergeSort(a []int) []int {
	// base case
	if len(a) <= 1 {
		return a
	}

	var left []int
	var right []int

	left, right = algoutils.Split(a)
	left = MergeSort(left)
	right = MergeSort(right)
	return Merge(left,right)
}

/*
	- while i < len(left) and j < len(right)
	4 scenarios
	1) Values in both left and right to disburse.
	2) Values in left are gone, right still has values.
	3) Values in right are gone, left still has values.
	4) Both left and right have been completely processed.

	There's room for some different ways of doing this.
	- Maybe reimplement using slices more effectively? left := left[1:]
*/
func Merge(left []int, right []int) []int {
	l,r,i := 0,0,0
	arr := make([]int, len(left) + len(right))

	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				arr[i] = left[l]
				l++
			} else {
				arr[i] = right[r]
				r++
			}
		} else if l >= len(left) {
			arr[i] = right[r]
			r++
		} else if r >= len(right) {
			arr[i] = left[l]
			l++
		}
		i++
	}
	return arr
}
