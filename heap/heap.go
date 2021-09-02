package heap

// Heap represents an interface of common methods between min and max heaps.
type Heap interface {
	// Push adds a new element to the heap.
	Push(value int)

	// Pop returns the top key, and removes it from the heap.
	Pop() int

	// Peek returns the firt element of the heap without removing it. It panics in case there is not items.
	Peek() int

	// Len returns the size of the heap.
	Len() int
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(parentIndex int) int {
	return parentIndex*2 + 1
}

func right(parentIndex int) int {
	return parentIndex*2 + 2
}
