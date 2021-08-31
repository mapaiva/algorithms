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

// AddRecursive will take in a word and insert it into the trie recursively.
func (t *Trie) AddRecursive(word string) {
	t.add(word, 0, t.root)
}

func (t *Trie) add(word string, index int, node *Node) {
	if len(word) == index {
		node.isEnd = true
		return
	}

	charIndex := word[index] - 'a'
	if node.children[charIndex] == nil {
		node.children[charIndex] = &Node{}
	}
	node.children[charIndex].size++

	t.add(word, index+1, node.children[charIndex])
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
