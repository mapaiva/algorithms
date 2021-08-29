package btree_test

import (
	"algorithms/btree"
	"testing"
)

func TestSum(t *testing.T) {
	root := &btree.Node{
		Data: 2,
		Left: &btree.Node{
			Data: 3,
			Left: &btree.Node{
				Data: 5,
			},
			Right: &btree.Node{
				Data: 6,
			},
		},
		Right: &btree.Node{
			Data: 4,
		},
	}

	if btree.Sum(root) != 20 {
		t.Errorf("expected %d got %d", 20, btree.Sum(root))
	}
}

func TestHeight(t *testing.T) {
	root := &btree.Node{Data: 3}
	btree.Add(root, 2)
	btree.Add(root, 5)
	btree.Add(root, 1)
	btree.Add(root, 4)
	btree.Add(root, 6)
	btree.Add(root, 7)

	if btree.Height(root) != 3 {
		t.Errorf("expected %d got %d", 3, btree.Height(root))
	}
}
