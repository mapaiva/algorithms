package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

// Len is the number of elements in the collection.
func (h MinHeap) Len() int { return len(h) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (h MinHeap) Less(i int, j int) bool { return h[i] < h[j] }

// Swap swaps the elements with indexes i and j.
func (h MinHeap) Swap(i int, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	oldLen := len(old)
	len := oldLen - 1
	extracted := old[len]
	*h = old[0:len]
	return extracted
}

func (h MinHeap) Peek() int {
	return h[0]
}

type MaxHeap []int

// Len is the number of elements in the collection.
func (h MaxHeap) Len() int {
	return len(h)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (h MaxHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

// Swap swaps the elements with indexes i and j.
func (h MaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	len := len(*h) - 1
	extracted := old[len]
	*h = old[0:len]
	return extracted
}

func (h MaxHeap) Peek() int {
	return h[0]
}

func main() {
	fmt.Println("Run go test -v ./cmd/heap")
}

func runningMedian(items []int) []float64 {
	highers := &MaxHeap{}
	lowers := &MinHeap{}

	var medians []float64
	for _, number := range items {
		if highers.Len() == 0 || highers.Peek() > number {
			heap.Push(highers, number)
		} else {
			heap.Push(lowers, number)
		}

		rebalance(highers, lowers)
		medians = append(medians, median(highers, lowers))
	}

	return medians
}

func rebalance(highers *MaxHeap, lowers *MinHeap) {
	if highers.Len() > lowers.Len()+1 {
		heap.Push(lowers, heap.Pop(highers))
	} else if lowers.Len() > highers.Len()+1 {
		heap.Push(highers, heap.Pop(lowers))
	}
}

func median(highers *MaxHeap, lowers *MinHeap) float64 {
	if highers.Len() > lowers.Len() {
		return float64(highers.Peek())
	}
	if lowers.Len() > highers.Len() {
		return float64(lowers.Peek())
	}

	return (float64(highers.Peek()) + float64(lowers.Peek())) / 2
}
