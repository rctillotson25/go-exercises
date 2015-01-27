/*

   This program will take in two strings and determine
   if they are anagrams of each other.

   Authored on 1/24/2015 by rctillotson25

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("This command-line utility determines if two strings are anagrams of each other.")

	s1 := getInput()
	s2 := getInput()

	fmt.Println("Are they anagrams? ", AreAnagrams(s1, s2))

}

func AreAnagrams(s1 string, s2 string) bool {

	if len(s1) == 0 && len(s2) == 0 {
		return true
	} else {
		// find location of the first rune in s1 in s2
		i := strings.IndexRune(s2, rune(s1[0]))

		// -1 = does not exist
		if i == -1 {
			return false
		} else {
			// s2[:i}+s2[i+1:] returns a new string excluding i
			return AreAnagrams(s1[1:], s2[:i]+s2[i+1:])
		}
	}

}

func getInput() string {
	console := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a string: ")

	input, err := console.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return input
}
