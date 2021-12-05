package array

type structure struct {
	items []int
}

func (s *structure) Add(elem int) {
	// @TODO: need to implement
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
	a := &structure{
		items: make([]int, len(elements)),
	}

	for i := range elements {
		a.items[i] = elements[i]
	}

	return a
}
