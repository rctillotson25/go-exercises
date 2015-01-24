/*

   This module holds the basic implementation of a an int node data
   structure.


   Node Type Attributes:

   Value interface{}
   next  *Node

   Authored 1/23/2015 by rctillotson25

*/

package structs

// need a wrapper class to pass
// compiler checks - no generics in Go.
type IntNode struct {
	intNode *Node
}

// Initialize new Int Node with given value
func InitIntNode(i int) *IntNode {
	node := &IntNode{new(Node)}
	node.SetValue(i)
	return node
}

func NewIntNode() *IntNode {
	return &IntNode{new(Node)}
}

func (n *IntNode) SetValue(v int) {
	n.intNode.Value = v
}

/*
	Cases:

	1) next is null (there is no next)
	2) next is not null, but does not have a next pointer
	3) next is not null, and there it does have a next pointer
	//TODO - FIX THIS ROUTINE - DOES NOT WORK PROPERLY

*/
func (n *IntNode) Next() *IntNode{
		return InitIntNode(n.intNode.next.Value.(int))
}

func (n *IntNode) Value() int {
	// prevent cast to int panic 
	if n.intNode.Value == nil {
		return (*new(int))
	} else {
		return n.intNode.Value.(int)
	}
}

func (n *IntNode) SetNext(i *IntNode) {
	n.intNode.next = i.intNode
}
