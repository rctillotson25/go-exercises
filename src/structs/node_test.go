package structs

import (
	"fmt"
	"testing"
)

func TestIntNode(t *testing.T) {

	n := NewIntNode()
	n.SetValue(7)
	n.SetNext(NewIntNode())
	s.intNode.next.SetValue(8)
	fmt.Println(n.intNode.Value)
	fmt.Println(n.intNode.next.intNode.value)
}
