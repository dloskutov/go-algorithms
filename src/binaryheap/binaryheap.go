package binaryheap

type heapType int

const (
	Min heapType = 1
	Max heapType = 2
)

type Element struct {
	Priority int
	Value    int
}

type binaryHeap struct {
	heapType heapType
	array    []*Element
}

func (h *binaryHeap) Insert(elem *Element) {
	// @TODO: need to implement
}

func (h *binaryHeap) Remove(elem *Element) bool {
	// @TODO: need to implement
	return false
}

func (h *binaryHeap) Update(elem *Element) bool {
	// @TODO: need to implement
	return false
}

func (h *binaryHeap) Top() *Element {
	// @TODO: need to implement
	return nil
}

func (h *binaryHeap) Peek() *Element {
	// @TODO: need to implement
	return nil
}

func New(elements []*Element, heapType heapType) *binaryHeap {
	h := &binaryHeap{
		heapType: heapType,
		array:    make([]*Element, len(elements)),
	}

	for i := range elements {
		h.Insert(elements[i])
	}

	return h
}
