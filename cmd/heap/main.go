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
		if highers.Size() == 0 || highers.Peek() > number {
			highers.Add(number)
		} else {
			lowers.Add(number)
		}

		rebalance(lowers, highers)
		medians = append(medians, median(lowers, highers))
	}

	return medians
}

func rebalance(lowers heap.Heap, highers heap.Heap) {
	if highers.Size() > lowers.Size()+1 {
		lowers.Add(highers.Extract())
	} else if lowers.Size() > highers.Size()+1 {
		highers.Add(lowers.Extract())
	}
}

func median(lowers heap.Heap, highers heap.Heap) float64 {
	var bigger, smaller heap.Heap

	if lowers.Size() > highers.Size() {
		bigger = lowers
		smaller = highers
	} else {
		bigger = highers
		smaller = lowers
	}

	if bigger.Size() == smaller.Size() {
		return (float64(bigger.Peek()) + float64(smaller.Peek())) / 2
	}

	return float64(bigger.Peek())
}
