package heap

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) parent(i int) int { return (i - 1) / 2 }
func (h *MaxHeap) left(i int) int   { return 2*i + 1 }
func (h *MaxHeap) right(i int) int  { return 2*i + 2 }

func (h *MaxHeap) Insert(v int) {
	h.data = append(h.data, v)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 && h.data[i] > h.data[h.parent(i)] {
		h.data[i], h.data[h.parent(i)] = h.data[h.parent(i)], h.data[i]
		i = h.parent(i)
	}
}

func (h *MaxHeap) Pop() int {
	if len(h.data) == 0 {
		panic("heap empty")
	}
	max := h.data[0]
	last := h.data[len(h.data)-1]
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)
	return max
}

func (h *MaxHeap) heapifyDown(i int) {
	largest := i
	l, r := h.left(i), h.right(i)

	if l < len(h.data) && h.data[l] > h.data[largest] {
		largest = l
	}
	if r < len(h.data) && h.data[r] > h.data[largest] {
		largest = r
	}
	if largest != i {
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		h.heapifyDown(largest)
	}
}
