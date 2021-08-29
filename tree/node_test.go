package tree_test

import (
	"algorithms/tree"
	"testing"
)

func BenchmarkNodeAssertion(b *testing.B) {
	var value interface{}
	value = tree.NewNode(42)

	for i := 0; i < b.N; i++ {
		_ = value.(*tree.Node)
	}
}
