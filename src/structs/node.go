package structs

type IntNode struct {
	intNode *Node
}
type Node struct {
	Value interface{}
	next  *Node
}

func NewIntNode() *IntNode {
	return &IntNode{new(Node)}
}

func (n *IntNode) SetValue(v int) {
	n.intNode.Value = v
}

func (n *IntNode) SetNext(i *IntNode) {
	n.intNode.next
}
