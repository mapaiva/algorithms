package heap

// MaxHeap represents a max heap data structure.
type MaxHeap struct {
	items []int
}

// NewMaxHeap returns a new heap.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Push adds a new element to the heap.
func (h *MaxHeap) Push(value int) {
	h.items = append(h.items, value)
	h.maxHeapifyUp(len(h.items) - 1)
}

// Pop returns the top key, and removes it from the heap.
func (h *MaxHeap) Pop() int {
	if len(h.items) == 0 {
		panic("empty heap")
	}

	extracted := h.items[0]
	l := h.Len() - 1
	h.items[0] = h.items[l]
	h.items = h.items[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// Peek returns the firt element of the heap without removing it. It panics in case there is not items.
func (h *MaxHeap) Peek() int {
	if len(h.items) == 0 {
		panic("empty heap")
	}
	return h.items[0]
}

// Len returns the size of the heap.
func (h *MaxHeap) Len() int {
	return len(h.items)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.items[parent(index)] < h.items[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := h.Len() - 1
	leftIndex, rightIndex := left(index), right(index)
	childToCompare := 0

	for leftIndex <= lastIndex {
		if leftIndex == lastIndex {
			childToCompare = leftIndex
		} else if h.items[leftIndex] > h.items[rightIndex] {
			childToCompare = leftIndex
		} else {
			childToCompare = rightIndex
		}

		if h.items[index] < h.items[childToCompare] {
			h.swap(index, childToCompare)
			index := childToCompare
			leftIndex, rightIndex = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}
