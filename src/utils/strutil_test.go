/*
   Testing class for the strutils.go class which implements 
   common string functions.

   Authored on 1/24/2015 by rctillotson25


*/
package utils


import (
	"testing"
	"fmt"
)


func TestCharCountMap(t *testing.T) {

	char_count := CharCountMap("Test String")

	for key, value := range char_count {
		fmt.Println("Char: ", string(key), "Count: ", value)
	}
}
