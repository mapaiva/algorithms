package tree_test

import (
	"algorithms/tree"
	"testing"
)

func TestSum(t *testing.T) {
	root := &tree.Node{
		Data: 2,
		Left: &tree.Node{
			Data: 3,
			Left: &tree.Node{
				Data: 5,
			},
			Right: &tree.Node{
				Data: 6,
			},
		},
		Right: &tree.Node{
			Data: 4,
		},
	}

	if tree.Sum(root) != 20 {
		t.Errorf("expected %d got %d", 20, tree.Sum(root))
	}
}

// func TestHeight(t *testing.T) {
// 	root := &tree.Node{
// 		Data: 2,
// 		Left: &tree.Node{
// 			Data: 3,
// 			Left: &tree.Node{
// 				Data: 5,
// 			},
// 			Right: &tree.Node{
// 				Data: 6,
// 			},
// 		},
// 		Right: &tree.Node{
// 			Data: 4,
// 		},
// 	}

// 	if tree.Height(root) != 3 {
// 		t.Errorf("expected %d got %d", 3, tree.Height(root))
// 	}
// }
