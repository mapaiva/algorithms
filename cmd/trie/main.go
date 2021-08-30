package main

import (
	"algorithms/trie"
	"fmt"
)

func main() {
	t := trie.New()
	t.Add("hack")
	t.Add("hackerank")

	fmt.Println(t.FindCount("hack"))
}
