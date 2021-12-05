package array

import "fmt"

var (
	ErrInvalidIndex = fmt.Errorf("invalid index")
)

type structure struct {
	items []int
}

func (s *structure) Add(value int) {
	s.items = append(s.items, value)
}

func (s *structure) Remove(index int) error {
	if index < 0 || index >= s.Size() {
		return ErrInvalidIndex
	}
	for index < s.Size()-1 {
		s.items[index] = s.items[index+1]
		index++
	}
	s.items = s.items[:len(s.items)-1]
	return nil
}

func (s *structure) Update(index int, value int) error {
	if index < 0 || index >= s.Size() {
		return ErrInvalidIndex
	}
	s.items[index] = value
	return nil
}

func (s *structure) Get(index int) (int, error) {
	if index < 0 || index >= s.Size() {
		return 0, ErrInvalidIndex
	}
	return s.items[index], nil
}

func (s *structure) Size() int {
	return len(s.items)
}

func New(elements []int) *structure {
	a := &structure{
		items: make([]int, len(elements)),
	}

	for i := range elements {
		a.items[i] = elements[i]
	}

	return a
}
