package heap

// Heap represents an interface of common methods between min and max heaps.
type Heap interface {
	// Add adds a new element to the heap.
	Add(value int)

	// Extract returns the top key, and removes it from the heap.
	Extract() int

	// Peek returns the firt element of the heap without removing it. It panics in case there is not items.
	Peek() int

	// Size returns the size of the heap.
	Size() int
}
