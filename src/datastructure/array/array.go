package array

import "fmt"

var (
	ErrInvalidIndex = fmt.Errorf("invalid index")
)

type Array struct {
	items []interface{}
}

func (s *Array) Add(value interface{}) {
	s.items = append(s.items, value)
}

func (s *Array) Remove(index int) error {
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

func (s *Array) Update(index int, value interface{}) error {
	if index < 0 || index >= s.Size() {
		return ErrInvalidIndex
	}
	s.items[index] = value
	return nil
}

func (s *Array) Get(index int) (interface{}, error) {
	if index < 0 || index >= s.Size() {
		return 0, ErrInvalidIndex
	}
	return s.items[index], nil
}

func (s *Array) Size() int {
	return len(s.items)
}

func New(elements []interface{}) *Array {
	a := &Array{
		items: make([]interface{}, len(elements)),
	}

	for i := range elements {
		a.items[i] = elements[i]
	}

	return a
}
