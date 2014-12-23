package main

import (
	"fmt"
	"random"
)

func main() {
	arr := random.RandIntN(100, 100)
	fmt.Println(len(arr))
	fmt.Println(arr)
	fmt.Println(search(0, len(arr), arr))
}

/*
-------------------------------
This function will return the index of integer n
if it exists in the array given at s.
-------------------------------
*/
func search(low int, high int, arr []int) int {
	if low > high {
		return -1
	}
	else if low
	return 1
}
