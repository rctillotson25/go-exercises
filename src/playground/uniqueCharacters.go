/*

   This command-line application gets the user input and determines
   if the string entered has all unique characters.

   Authored 1/17/2015  by Ryan Tillotson

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input := getInput()

	result := UniqueString(input)

	fmt.Println("Result: ", result)

}

/*
   Base case:
   	Empty string is unique and indicates all characters unique
   Failure case:
        Tail contains the head of the string
   Recursive case:
        Check the tail of the string

*/
func UniqueString(s string) bool {
	if len(s) <= 0 {
		return true
	} else if strings.Contains(string(s[1:]), string(s[0])) == true {
		return false
	} else {
		return UniqueString(s[1:])
	}
}

func getInput() string {
	console := bufio.NewReader(os.Stdin)
	fmt.Println("This command-line utility determines if a string has unique characters.")
	fmt.Println("Please enter a string: ")

	input, err := console.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return input
}
