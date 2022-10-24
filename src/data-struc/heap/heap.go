package heap

func gt(lhs int, rhs int) bool {
	return lhs > rhs
}

func lt(lhs int, rhs int) bool {
	return lhs < rhs
}

func NewMaxHeap() Heap {
	return Heap{comp: lt}
}

func NewMinHeap() Heap {
	return Heap{comp: gt}
}

type Heap struct {
	data []int
	comp func(int, int) bool
}

func (h* Heap) Size() int {
	return len(h.data)
}

func (h* Heap) Insert(value int) {
	h.data = append(h.data, value)
	h.maxHeapifyUp(h.last())
}

func (h* Heap) Pop() int {
	value := h.data[0]

	h.data[0] = h.data[h.last()]
	h.data = h.data[:h.last()]

	h.maxHeapifyDown(0)

	return value
}

func parent(idx int) int {
	return (idx-1)/2
}

func left(idx int) int {
	return idx*2+1
}

func right(idx int) int {
	return idx*2+2
}

func (h* Heap) swap(lhs, rhs int)  {
	h.data[lhs], h.data[rhs] = h.data[rhs], h.data[lhs]
}

func (h* Heap) last() int {
	if len(h.data) == 0 {
		return 0
	}
	return len(h.data) - 1
}

func (h* Heap) maxHeapifyUp(idx int) {
	for h.comp(h.data[parent(idx)], h.data[idx]) {
		h.swap(parent(idx), idx)
		idx = parent(idx)
	}
}

func (h* Heap) maxHeapifyDown(idx int) {
	node := h.last()


	for left(idx) <= h.last() {
		if left(idx) == h.last() {
			node = left(idx)
		} else if ! h.comp(h.data[left(idx)], h.data[right(idx)]) {
			node = left(idx)
		} else {
			node = right(idx)
		}

		if h.comp(h.data[idx], h.data[node]) {
			h.swap(idx, node)
			idx = node
		} else {
			return
		}
	}
}
