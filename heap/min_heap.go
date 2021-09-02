package heap

// MinHeap represents a max heap data structure.
type MinHeap struct {
	items []int
}

// NewMinHeap returns a new heap.
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// Add adds a new element to the heap.
func (h *MinHeap) Add(value int) {
	h.items = append(h.items, value)
	h.minHeapifyUp(len(h.items) - 1)
}

// Extract returns the top key, and removes it from the heap.
func (h *MinHeap) Extract() int {
	if len(h.items) == 0 {
		panic("empty heap")
	}

	extracted := h.items[0]
	l := h.Size() - 1
	h.items[0] = h.items[l]
	h.items = h.items[:l]

	h.minHeapifyDown(0)

	return extracted
}

// Peek returns the firt element of the heap without removing it. It panics in case there is not items.
func (h *MinHeap) Peek() int {
	if len(h.items) == 0 {
		panic("empty heap")
	}
	return h.items[0]
}

// Size returns the size of the heap.
func (h *MinHeap) Size() int {
	return len(h.items)
}

func (h *MinHeap) minHeapifyUp(index int) {
	for h.items[parent(index)] > h.items[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) minHeapifyDown(index int) {
	lastIndex := h.Size() - 1
	leftIndex, rightIndex := left(index), right(index)
	childToCompare := 0

	for leftIndex <= lastIndex {
		if leftIndex == lastIndex {
			childToCompare = leftIndex
		} else if h.items[leftIndex] < h.items[rightIndex] {
			childToCompare = leftIndex
		} else {
			childToCompare = rightIndex
		}

		if h.items[index] > h.items[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			leftIndex, rightIndex = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}
