package main

import (
	"fmt"
)

/* Practice with pointers and dereferencing variables.

   This example shows how pointers work in Go at the base
   level and illustrates the differences between passing
   by value and passing by reference.

*/
func main() {
	s := 1
	fmt.Println(s)

	// s passed by reference
	// will be incremented by 1
	reference(&s)
	fmt.Println(s)

	// s passed by value
	// value of s constant
	value(s)
	fmt.Println(s)

}

/*
   x is passed by reference.
   The memory location of the
   variable from the main function
   is passed in. Calling *x dereferences
   this variable and adds 1 to it.
   The variable in the main function
   will be modified.
*/
func reference(x *int) {
	*x += 1
}

/*
   x is passed by value. Thus,
   since this function isn't looking
   at the original memory location,
   the parameter will not be modified
   in its original scope.
*/
func value(x int) {
	x += 11
}
