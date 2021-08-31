package main

import (
	"algorithms/heap"
	"fmt"
)

func main() {
	m := heap.NewMaxHeap()

	// items := []int{3, 4, 8, 9, 7, 10, 9, 15, 20, 13}
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range items {
		m.Add(i)
		fmt.Println(m)
	}

}
