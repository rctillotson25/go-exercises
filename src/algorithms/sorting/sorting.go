/*

    This module contains several types of sorting algorithms and their space / time complexities.

    Current Status of Algorithms:

    Insertion Sort - Complete
    Bubble Sort    - Not Started
    Merge Sort     - Complete
    Quick Sort     - Not Started
    Heap Sort      - Not Started


    To initialize tests of the algorithms, add the algorithm to sorting_test.go.

*/



package sorting

import (
	"utils"
)

/*
        ---------------------------
        Insertion Sorting Algorithm
        ---------------------------

	* Space complexity: O(n) total - O(1) auxiliary. 
	In place sorting, so no additional space is necessary.
	* Best-case time: O(n) - occurs when the array is already sorted.
	* Average-case time: O(n^2)
	* Worst-case time: O(n^2)

*/
func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				utils.Swap(a, j, j-1)
			}
		}
	}
	return a
}

/*
	---------------------------
	Merge Sorting Algorithm
	---------------------------

	* Space complexity: O(n) total - O(n) auxiliary.
	* Best-case time: O(nlog(n)) -
	* Average-case time: O(nlog(n))
	* Worst-case time: O(nlog(n))
*/
func MergeSort(a []int) []int {

	// base case
	if len(a) <= 1 {
		return a
	}

	var left []int
	var right []int

	left, right = utils.Split(a)
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left,right)
}

/*
	Merge and return 2 slices of integers.

	4 key scenarios
	1) Values in both left and right to disburse.
	2) Values in left are gone, right still has values.
	3) Values in right are gone, left still has values.
	4) Both left and right have been completely processed.
*/
func merge(left []int, right []int) []int {
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

