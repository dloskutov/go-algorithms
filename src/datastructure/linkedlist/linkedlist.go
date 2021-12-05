package linkedlist

type node struct {
	value int
	next  *node
}

type structure struct {
	head *node
}

func (s *structure) Add(value int) {
	n := &node{
		value: value,
	}
	if s.head == nil {
		s.head = n
		return
	}
	nextNode := s.head
	for nextNode.next != nil {
		nextNode = nextNode.next
	}
	nextNode.next = n
}

func (s *structure) Remove(index int) error {
	// @TODO: need to implement
	return nil
}

func (s *structure) Update(index int, elem int) error {
	// @TODO: need to implement
	return nil
}

func (s *structure) Get(index int) (int, error) {
	// @TODO: need to implement
	return 0, nil
}

func New(elements []int) *structure {
	s := &structure{}

	for i := range elements {
		s.Add(elements[i])
	}

	return s
}
