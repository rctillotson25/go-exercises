package main

import (
	"fmt"
	"utils"
	"os"
)


// TODO - Add ability to get strings with spaces in them.

func main() {

	s := getInput()

	fmt.Println("String: \t\t",  s)
	fmt.Println("Reversed String: \t", reverseString(s))

	os.Exit(0)
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

func getInput() string {
	var s string

	fmt.Print("Enter a string to be reversed: ")
	_, err := fmt.Scanf("%s", &s)

	if err != nil {
		fmt.Println("There was an error in your input.")
		os.Exit(1)
	}

	return s
}
