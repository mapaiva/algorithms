package tree

import (
	"container/list"
	"fmt"
)

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

// TraversalInOrder returns an int array with the result of a in-order traversal over the tree.
func TraversalInOrder(root *Node) []int {
	panic("TODO")
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
