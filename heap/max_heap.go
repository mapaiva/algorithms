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

	n := h.Len() - 1
	h.swap(0, n)
	h.maxHeapifyDown(0)

	extracted := h.items[n]
	h.items = h.items[:n]
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
	last := h.Len() - 1

	for {
		l, r := left(index), right(index)
		if l >= last || l < 0 { // l < 0 after int overflow
			break
		}
		childToCompare := l
		if r < last && h.items[r] > h.items[l] {
			childToCompare = r
		}
		if h.items[childToCompare] < h.items[index] {
			break
		}
		h.swap(childToCompare, index)
		index = childToCompare
	}
}

func (h *MaxHeap) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}
