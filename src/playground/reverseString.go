package main

import (
	"fmt"
	"utils"
)


// TODO - Add user input for the reverse string class.

func main() {
	s := []string{"Apple", "Banana", "Carrot"}

	for i := range s {
		fmt.Println(i, "- String: \t\t",  s[i])
		fmt.Println(i, "- Reversed String: \t", reverseString(s[i]))
	}
}

func reverseString(s string) string {
	var low int = 0
	var high int = len(s) - 1
	var strArr []byte

	strArr = make([]byte, len(s))
	copy(strArr, s)

	for low <= high {
		utils.SwapBytes(&strArr, low, high)
		low += 1
		high -= 1
	}
	return string(strArr)
}


