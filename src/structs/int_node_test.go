/*

   Testing package for basic int node implementation in golang.
   Ensures that linking nodes works properly and validates the results.

   Authored on 1/22/2015 by rctillotson25


*/

package structs

import (
	"fmt"
	"testing"
)


func TestIntNode(t *testing.T) {

	i1 := NewIntNode()
	i1.SetValue(1)

	i2 := NewIntNode()
	i2.SetValue(2)

	i1.SetNext(i2)

	fmt.Println(i1.Value())
	fmt.Println(i1.Next().Value())
	
	if i1.Value() != 1 || i1.Next().Value() != 2 {
		t.Errorf("Incorrect values for nodes i1: ", i1.Value(), " i2: ", i2.Next().Value())
	}

}

func TestNull(t *testing.T) {

	// test null object
	i := new(IntNode)
	fmt.Println(i)

	// test null values
	x := NewIntNode()
	_ = x.Value()
	_ = x.intNode

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Null values caused the program to panic.")

		}
	}()

}
