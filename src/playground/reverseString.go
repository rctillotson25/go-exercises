package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []string{"Apple", "Banana", "Carrot"}

	for i := range s {
		fmt.Println("String " + strconv.Itoa(i) + ": " + s[i])
	}
}

func reverseString(s string) string {
	revstr := make([]byte, len(s))

	fmt.Println(revstr)
	return s
}


