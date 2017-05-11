package heap

type Item interface {
	Less(than interface{}) bool
}

type Heap struct {
	data []Item
	min  bool
}

func New(isMin bool) *Heap {
	return &Heap{
		data: []Item{nil},
		min:  isMin,
	}
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 1
}

func (h *Heap) Len() int {
	return len(h.data) - 1
}

func (h *Heap) Less(i, j int) bool {
	if h.min {
		return h.data[i].Less(h.data[j])
	}
	return h.data[j].Less(h.data[i])

}

func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) Swim(k int) {
	for k > 1 && h.Less(k, k/2) {
		h.Swap(k/2, k)
		k = k / 2
	}
}

func (h *Heap) Sink(k int) {
	n := h.Len()
	for 2*k <= n {
		j := 2 * k
		if j < n && h.Less(j+1, j) {
			j++
		}
		if !h.Less(j, k) {
			break
		}
		h.Swap(k, j)
		k = j
	}
}

func (h *Heap) Insert(n Item) {
	h.data = append(h.data, n)
	h.Swim(h.Len())
}

func (h *Heap) Extract() (n Item) {
	if h.IsEmpty() {
		return
	}

	n = h.data[1]
	h.Swap(1, h.Len())
	h.data = h.data[0:h.Len()]

	if h.Len() > 1 {
		h.Sink(1)
	}
	return
}

func (h *Heap) Top() (n Item) {
	if h.IsEmpty() {
		return
	}

	return h.data[1]
}
