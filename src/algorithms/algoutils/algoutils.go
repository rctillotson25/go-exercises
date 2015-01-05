package algoutils;


import (
)

// Swap the two integers located at indices i and j
func Swap(A []int, i int, j int) {
	t := A[i]
	A[i] = A[j]
	A[j] = t
}
