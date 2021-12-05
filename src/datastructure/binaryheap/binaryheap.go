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

type structure struct {
	heapType heapType
	array    []*Element
}

func (s *structure) Insert(elem *Element) {
	// @TODO: need to implement
}

func (s *structure) Remove(elem *Element) error {
	// @TODO: need to implement
	return nil
}

func (s *structure) Update(elem *Element) error {
	// @TODO: need to implement
	return nil
}

func (s *structure) Top() *Element {
	// @TODO: need to implement
	return nil
}

func (s *structure) Peek() *Element {
	// @TODO: need to implement
	return nil
}

func New(elements []*Element, heapType heapType) *structure {
	s := &structure{
		heapType: heapType,
		array:    make([]*Element, len(elements)),
	}

	for i := range elements {
		s.Insert(elements[i])
	}

	return s
}
