package random

import (
	"math/rand"
	"time"
)

/*
	Provide n - the number of random numbers between 0
	and the upper bound.

*/
func RandIntN(n int, upper int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(upper)
	}
	return arr
}
