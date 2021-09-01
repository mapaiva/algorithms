package main

import (
	"algorithms/heap"
	"fmt"
)

func main() {
	var medians []float64
	lowers := heap.NewMinHeap()
	highers := heap.NewMaxHeap()

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range items {
		lowers.Add(number)
		highers.Add(number)

		medians = append(medians, median(lowers, highers))
	}

	fmt.Println(lowers)
	fmt.Println(highers)
	fmt.Println(medians)
}

func median(lowers *heap.MinHeap, highers *heap.MaxHeap) float64 {
	var biggerHeap, smallerHeap heap.Heap

	if lowers.Size() > highers.Size() {
		biggerHeap = lowers
		smallerHeap = highers
	} else {
		biggerHeap = highers
		smallerHeap = lowers
	}

	if biggerHeap.Size() == smallerHeap.Size() {
		return (float64(biggerHeap.Peek()) + float64(smallerHeap.Peek())) / 2
	}

	return float64(biggerHeap.Peek())
}
