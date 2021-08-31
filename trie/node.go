package trie

// alphabetSize represents the quantity of letters in the alphabet.
const alphabetSize = 26

// Node represents a trie node.
type Node struct {
	children [alphabetSize]*Node
	isEnd    bool
	size     int
}
