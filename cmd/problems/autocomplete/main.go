package main

import (
	"fmt"
	"sort"
)

const (
	Limit = 3
)

type Trie struct {
	Root *Node
}

func (t *Trie) Put(word string) {
	current := t.Root
	for _, char := range word {
		if current.Children[char] == nil {
			current.Children[char] = &Node{Children: make(map[rune]*Node)}
		}
		current = current.Children[char]
	}
	current.IsWord = true
}

func (t *Trie) Find(suffix string) *Node {
	current := t.Root
	for _, char := range suffix {
		if current.Children[char] == nil {
			return nil
		}
		current = current.Children[char]
	}

	return current
}

func (t *Trie) Search(suffix string) []string {
	node := t.Find(suffix) // O(s) onde s = len(suffix)
	if node == nil {
		return nil
	}

	return node.FindSuffixes(suffix)
}

type Node struct {
	Children map[rune]*Node
	IsWord   bool
}

func (n *Node) FindSuffixes(suffix string) []string {
	var results []string
	if n.IsWord {
		results = append(results, suffix)
	}

	if n.Children != nil {
		n.findSuffixes(suffix, &results)
	}

	return results
}

func (n *Node) findSuffixes(suffix string, results *[]string) {
	var keys []int
	for char := range n.Children { // O(n) onde n = len(children)
		keys = append(keys, int(char))
	}

	sort.Ints(keys) // O(n*logn) sort performs a quicksort

	for _, key := range keys { // O(n)
		char := rune(key)
		child := n.Children[char]
		w := suffix + string(char)

		if child.IsWord {
			*results = append(*results, w)
		}

		if len(*results) == Limit {
			break
		}

		child.findSuffixes(w, results) // O(n)
	}
}

func autocomplete(repository []string, customerQuery string) [][]string {
	trie := &Trie{Root: &Node{Children: make(map[rune]*Node)}}
	for _, word := range repository { // O(r) onde r = len(repository)
		trie.Put(word) // O(w) onde w = len(word)
	}

	var results [][]string
	for i := 2; i <= len(customerQuery); i++ { // O(c-2) => O(c) onde c = len(customerQuery)
		subQuery := customerQuery[0:i] // O(1)

		results = append(results, trie.Search(subQuery))
	}

	return results
}

func main() {
	repository := []string{
		// "mobile",
		// "mouse",
		// "moneypot",
		// "monitor",
		// "mousepad",
		"code",
		"codePhone",
		"coddle",
		"coddles",
		"codes",
		"coddle",
	}
	// customerQuery := "mouse"
	customerQuery := "code"

	for _, result := range autocomplete(repository, customerQuery) {
		fmt.Println(result)
	}
}
