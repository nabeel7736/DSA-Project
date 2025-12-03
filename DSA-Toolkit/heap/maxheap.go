package heap

type MaxHeap struct {
	Data []int
}

func (h *MaxHeap) parent(i int) int { return (i - 1) / 2 }
func (h *MaxHeap) left(i int) int   { return 2*i + 1 }
func (h *MaxHeap) right(i int) int  { return 2*i + 2 }

func (h *MaxHeap) Insert(v int) {
	h.Data = append(h.Data, v)
	h.heapifyUp(len(h.Data) - 1)
}

func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 && h.Data[i] > h.Data[h.parent(i)] {
		h.Data[i], h.Data[h.parent(i)] = h.Data[h.parent(i)], h.Data[i]
		i = h.parent(i)
	}
}

func (h *MaxHeap) Pop() int {
	if len(h.Data) == 0 {
		panic("heap empty")
	}
	max := h.Data[0]
	last := h.Data[len(h.Data)-1]
	h.Data[0] = last
	h.Data = h.Data[:len(h.Data)-1]
	h.heapifyDown(0)
	return max
}

func (h *MaxHeap) heapifyDown(i int) {
	largest := i
	l, r := h.left(i), h.right(i)

	if l < len(h.Data) && h.Data[l] > h.Data[largest] {
		largest = l
	}
	if r < len(h.Data) && h.Data[r] > h.Data[largest] {
		largest = r
	}
	if largest != i {
		h.Data[i], h.Data[largest] = h.Data[largest], h.Data[i]
		h.heapifyDown(largest)
	}
}
