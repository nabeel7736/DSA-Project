package heap

type MinHeap struct {
	data []int
}

func (h *MinHeap) parent(i int) int { return (i - 1) / 2 }
func (h *MinHeap) left(i int) int   { return 2*i + 1 }
func (h *MinHeap) right(i int) int  { return 2*i + 2 }

func (h *MinHeap) Insert(v int) {
	h.data = append(h.data, v)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) heapifyUp(i int) {
	for i > 0 && h.data[i] < h.data[h.parent(i)] {
		h.data[i], h.data[h.parent(i)] = h.data[h.parent(i)], h.data[i]
		i = h.parent(i)
	}
}

func (h *MinHeap) Pop() int {
	if len(h.data) == 0 {
		panic("heap empty")
	}
	min := h.data[0]
	last := h.data[len(h.data)-1]
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)
	return min
}

func (h *MinHeap) heapifyDown(i int) {
	smallest := i
	l, r := h.left(i), h.right(i)

	if l < len(h.data) && h.data[l] < h.data[smallest] {
		smallest = l
	}
	if r < len(h.data) && h.data[r] < h.data[smallest] {
		smallest = r
	}
	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.heapifyDown(smallest)
	}
}
