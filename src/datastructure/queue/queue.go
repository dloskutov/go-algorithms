package queue

type structure struct{}

func (s *structure) Enqueue(elem int) {
	// @TODO: need to implement
}

func (s *structure) Dequeue() (int, error) {
	// @TODO: need to implement
	return 0, nil
}

func New(elements []int) *structure {
	s := &structure{}

	return s
}
