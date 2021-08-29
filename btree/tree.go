package btree

import (
	"container/list"
	"fmt"
)

// BinaryTree represents a binary search tree.
type BinaryTree struct {
	root *Node
}

// NewBinaryTree creates a new binary tree.
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// Sum returns the sum of all tree nodes.
func (t *BinaryTree) Sum() int {
	if t.root == nil {
		return 0
	}

	stack := list.New()
	stack.PushFront(t.root)

	var sum int
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		node := e.Value.(*Node)

		sum += node.Data

		if node.Left != nil {
			stack.PushFront(node.Left)
		}
		if node.Right != nil {
			stack.PushFront(node.Right)
		}
	}
	return sum
}

// Add adds a new node on the binary search tree.
func (t *BinaryTree) Add(data int) {
	if t.root == nil {
		t.root = NewNode(data)
	} else {
		current := t.root
		for {
			if data < current.Data {
				if current.Left == nil {
					current.Left = NewNode(data)
					break
				} else {
					current = current.Left
				}
			}

			if data > current.Data {
				if current.Right == nil {
					current.Right = NewNode(data)
					break
				} else {
					current = current.Right
				}
			}
		}
	}
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
		return -1
	}

	leftHeight := Height(root.Left)
	rightHeight := Height(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1

}

// PrintInOrder prints the tree in-order: Left, Root, Right.
func PrintInOrder(root *Node) {
	if root.Left != nil {
		PrintInOrder(root.Left)
	}
	fmt.Printf(" %d ", root.Data)
	if root.Right != nil {
		PrintInOrder(root.Right)
	}
}

// PrintPreOrder prints the tree in pre-order: Root, Left, Right.
func PrintPreOrder(root *Node) {
	fmt.Printf(" %d ", root.Data)
	if root.Left != nil {
		PrintPreOrder(root.Left)
	}
	if root.Right != nil {
		PrintPreOrder(root.Right)
	}
}

// PrintPostOrder prints the tree in post-order: Left, Right, Root.
func PrintPostOrder(root *Node) {
	if root.Left != nil {
		PrintPreOrder(root.Left)
	}
	if root.Right != nil {
		PrintPreOrder(root.Right)
	}
	fmt.Printf(" %d ", root.Data)
}

// PrintLevelOrder prints the tree in level order.
func PrintLevelOrder(root *Node) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		v := e.Value
		root = v.(*Node)

		fmt.Printf(" %d ", root.Data)

		if root.Left != nil {
			queue.PushBack(root.Left)
		}
		if root.Right != nil {
			queue.PushBack(root.Right)
		}
	}
}
