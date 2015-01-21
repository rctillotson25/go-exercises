/*

   This command-line utility takes a line of input from the console
   and prints out the reversed version of the string.

   Created 1/15/2015 by rctillotson25


*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"utils"
)

// TODO - Improve the trimming of strings so that it doesn't pick up the \n at the end.

func main() {

	s := getInput()

	fmt.Println("String: \t\t", s)
	fmt.Println("Reversed String: \t", reverseString(s))

	os.Exit(0)
}

func reverseString(s string) string {
	var low int = 0
	var high int = len(s) - 1
	var strArr []byte

	// strings are read-only slices in Go - must copy string to swap
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

	console := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string to be reversed: ")
	s, err := console.ReadString('\n')

	if err != nil {
		fmt.Println("There was an error in your input.")
		os.Exit(1)
	}

	return s
}
