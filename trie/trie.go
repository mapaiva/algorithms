package trie

// Trie represents a trie.
type Trie struct {
	root *Node
}

// New creates a new trie.
func New() *Trie {
	return &Trie{root: &Node{}}
}

// Add will take in a word and insert it into the trie.
func (t *Trie) Add(word string) {
	wordLen := len(word)
	currentNode := t.root

	for i := 0; i < wordLen; i++ {
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}

		currentNode = currentNode.children[charIndex]
		currentNode.size++
	}

	currentNode.isEnd = true
}

// Find will take in a word and return true if that word is in the trie or false if it is not.
func (t *Trie) Find(word string) bool {
	wordLen := len(word)
	currentNode := t.root

	for i := 0; i < wordLen; i++ {
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}

func (t *Trie) FindCount(word string) int {
	wordLen := len(word)
	currentNode := t.root

	for i := 0; i < wordLen; i++ {
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			return 0
		}

		currentNode = currentNode.children[charIndex]
	}

	return currentNode.size
}
