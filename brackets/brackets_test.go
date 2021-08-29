package brackets_test

import (
	"algorithms/brackets"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tt := []struct {
		s        string
		expected bool
	}{
		{"{[()]}", true},
		{"{[(]}", false},
		{"i}][}}(}][))]", false},
		{"[](){()}", true},
		{"({}([][]))[]()", true},
		{"{)[](}]}]}))}(())(", false},
		{"{}(", false},
	}

	for _, test := range tt {
		if brackets.IsBalanced(test.s) != test.expected {
			t.Errorf("expected %s is balanced to be %v", test.s, test.expected)
		}
	}
}
