package heap

type MinHeap struct {
	Data []int
}

func (h *MinHeap) parent(i int) int { return (i - 1) / 2 }
func (h *MinHeap) left(i int) int   { return 2*i + 1 }
func (h *MinHeap) right(i int) int  { return 2*i + 2 }

func (h *MinHeap) Insert(v int) {
	h.Data = append(h.Data, v)
	h.heapifyUp(len(h.Data) - 1)
}

func (h *MinHeap) heapifyUp(i int) {
	for i > 0 && h.Data[i] < h.Data[h.parent(i)] {
		h.Data[i], h.Data[h.parent(i)] = h.Data[h.parent(i)], h.Data[i]
		i = h.parent(i)
	}
}

func (h *MinHeap) Pop() int {
	if len(h.Data) == 0 {
		panic("heap empty")
	}
	min := h.Data[0]
	last := h.Data[len(h.Data)-1]
	h.Data[0] = last
	h.Data = h.Data[:len(h.Data)-1]
	h.heapifyDown(0)
	return min
}

func (h *MinHeap) heapifyDown(i int) {
	smallest := i
	l, r := h.left(i), h.right(i)

	if l < len(h.Data) && h.Data[l] < h.Data[smallest] {
		smallest = l
	}
	if r < len(h.Data) && h.Data[r] < h.Data[smallest] {
		smallest = r
	}
	if smallest != i {
		h.Data[i], h.Data[smallest] = h.Data[smallest], h.Data[i]
		h.heapifyDown(smallest)
	}
}
