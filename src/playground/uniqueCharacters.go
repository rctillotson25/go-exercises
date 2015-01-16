package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	console := bufio.NewReader(os.Stdin)
	fmt.Println("This command-line utility determines if a string has unique characters.")
	fmt.Println("Please enter a string: ")

	input, err := console.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := UniqueString(input)

	fmt.Println("Result: ", result)

}

func UniqueString(s string) bool {
	if len(s) <= 0 {
		return true
	} else if strings.Contains(string(s[1:]), string(s[0])) == true {
		return false
	} else {
		return UniqueString(s[1:])
	}
}
