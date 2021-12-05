package stack

import "fmt"

var ErrEmpty = fmt.Errorf("stack is empty")

type structure struct {
	items []int
}

func (s *structure) Push(value int) {
	s.items = append(s.items, value)
}

func (s *structure) Pop() (int, error) {
	if s.Size() == 0 {
		return 0, ErrEmpty
	}
	value := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
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
