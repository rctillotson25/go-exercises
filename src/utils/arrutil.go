package utils

import ()

// Swap the two integers located at indices i and j
func Swap(a []int, i int, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

// Split a slice/array into two at the midpoint
// Return left and right slices.
func Split(a []int) ([]int, []int) {
	mid := len(a) / 2
	left := a[:mid]
	right := a[mid:]

	return left, right
}

// Return true if two slices of integers
// are identical (same order, length, values)
func Equal(a []int, b []int) bool {
	
	if len(a) != len(b) {
		return false
	} else {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
	}
	return true
}
