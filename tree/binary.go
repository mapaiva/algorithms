package tree

// BinaryTree represents an implementation of a binary tree.
type BinaryTree struct {
	Root *Node
}

// Add adds a new node onto the tree using binary search tree strategy.
func Add(root *Node, data int) {
	if data < root.Data {
		if root.Left == nil {
			root.Left = NewNode(data)
		} else {
			Add(root.Left, data)
		}
	}

	if data > root.Data {
		if root.Right == nil {
			root.Right = NewNode(data)
		} else {
			Add(root.Right, data)
		}
	}
}

// Sum returns the sum of all tree nodes.
func Sum(root *Node) int {
	if root == nil {
		return 0
	}

	return root.Data + Sum(root.Left) + Sum(root.Right)
}

// Height calculates the height of a tree.
func Height(root *Node) int {
	if root == nil {
		return 0
	}

	// return root.Data + Sum(root.Left) + Sum(root.Right)
	return 0
}
