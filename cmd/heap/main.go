package main

import (
	"algorithms/heap"
	"fmt"
)

func main() {
	items := []int{
		94455,
		20555,
		20535,
		53125,
		73634,
		148,
		63772,
		17738,
		62995,
		13401,
		95912,
		13449,
		92211,
		17073,
		69230,
		22016,
		22120,
		78563,
		16571,
	}

	fmt.Println(runningMedian(items))
}

func runningMedian(items []int) []float64 {
	var medians []float64
	lowers := heap.NewMinHeap()
	highers := heap.NewMaxHeap()

	for _, number := range items {
		if highers.Len() == 0 || highers.Peek() > number {
			highers.Push(number)
		} else {
			lowers.Push(number)
		}

		rebalance(lowers, highers)
		medians = append(medians, median(lowers, highers))
	}

	return medians
}

func rebalance(lowers heap.Heap, highers heap.Heap) {
	if highers.Len() > lowers.Len()+1 {
		lowers.Push(highers.Pop())
	} else if lowers.Len() > highers.Len()+1 {
		highers.Push(lowers.Pop())
	}
}

func median(lowers heap.Heap, highers heap.Heap) float64 {
	var bigger, smaller heap.Heap

	if lowers.Len() > highers.Len() {
		bigger = lowers
		smaller = highers
	} else {
		bigger = highers
		smaller = lowers
	}

	if bigger.Len() == smaller.Len() {
		return (float64(bigger.Peek()) + float64(smaller.Peek())) / 2
	}

	return float64(bigger.Peek())
}
