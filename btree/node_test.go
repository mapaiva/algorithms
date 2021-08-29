package btree_test

import (
	"algorithms/btree"
	"testing"
)

func BenchmarkNodeAssertion(b *testing.B) {
	var value interface{}
	value = btree.NewNode(42)

	for i := 0; i < b.N; i++ {
		_ = value.(*btree.Node)
	}
}
