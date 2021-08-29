package tree

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
