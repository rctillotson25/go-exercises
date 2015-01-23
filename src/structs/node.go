/*

   This module holds the basic Node type for the data structure.
   Need to implement different kinds of nodes to place values in them.

   Interface{} is more of a placeholder in this context.

   Authored 1/23/2015 by rctillotson25

*/

package structs

type Node struct {
	// every object extends interface{}
	Value interface{}
	next  *Node
}
