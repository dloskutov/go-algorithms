package queue

import "fmt"

var ErrEmpty = fmt.Errorf("queue is empty")

type structure struct {
	items []int
}

func (s *structure) Enqueue(value int) {
	s.items = append(s.items, value)
}

func (s *structure) Dequeue() (int, error) {
	if s.Size() == 0 {
		return 0, ErrEmpty
	}
	value := s.items[0]
	s.items = s.items[1:]
	return value, nil
}

func (s *structure) Size() int {
	return len(s.items)
}

func New(elements []int) *structure {
	s := &structure{
		items: make([]int, len(elements)),
	}

	for i := range elements {
		s.items[i] = elements[i]
	}

	return s
}
