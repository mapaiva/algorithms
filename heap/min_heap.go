package heap

// MinHeap represents a max heap data structure.
type MinHeap struct {
	items []int
}

// NewMinHeap returns a new heap.
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// Push adds a new element to the heap.
func (h *MinHeap) Push(value int) {
	h.items = append(h.items, value)
	h.minHeapifyUp(len(h.items) - 1)
}

// Pop returns the top key, and removes it from the heap.
func (h *MinHeap) Pop() int {
	if len(h.items) == 0 {
		panic("empty heap")
	}

	n := h.Len() - 1
	h.swap(0, n)
	h.minHeapifyDown(0)

	extracted := h.items[n]
	h.items = h.items[:n]
	return extracted
}

// Peek returns the firt element of the heap without removing it. It panics in case there is not items.
func (h *MinHeap) Peek() int {
	if len(h.items) == 0 {
		panic("empty heap")
	}
	return h.items[0]
}

// Len returns the size of the heap.
func (h *MinHeap) Len() int {
	return len(h.items)
}

func (h *MinHeap) minHeapifyUp(index int) {
	for h.items[parent(index)] > h.items[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) minHeapifyDown(index int) {
	last := h.Len() - 1

	for {
		l, r := left(index), right(index)
		if l >= last || l < 0 { // l < 0 after int overflow
			break
		}
		childToCompare := l
		if r < last && h.items[r] < h.items[l] {
			childToCompare = r
		}
		if h.items[childToCompare] > h.items[index] {
			break
		}
		h.swap(childToCompare, index)
		index = childToCompare
	}
}

func (h *MinHeap) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}
