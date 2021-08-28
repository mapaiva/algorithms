package tree

import "fmt"

// Node represents a binary tree node.
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// NewNode creates a new node.
func NewNode(data int) *Node {
	return &Node{Data: data}
}

func (n *Node) String() string {
	var left, right int

	if n.Left != nil {
		left = n.Left.Data
	}

	if n.Right != nil {
		right = n.Right.Data
	}

	return fmt.Sprintf("Data: %d, Left: %v, Right: %v", n.Data, left, right)
}
