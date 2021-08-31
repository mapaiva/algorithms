package heap

// MaxHeap represents a max heap data structure.
type MaxHeap struct {
	items []int
}

// NewMaxHeap returns a new heap.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Add adds a new element to the heap.
func (h *MaxHeap) Add(value int) {
	h.items = append(h.items, value)
	h.maxHeapifyUp(len(h.items) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.items[parent(index)] < h.items[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
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

func (h *MaxHeap) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}
