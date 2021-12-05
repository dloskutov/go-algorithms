package hashtable

type structure struct{}

func (s *structure) Put(key string, value int) {
	// @TODO: need to implement
}

func (s *structure) Remove(key string) error {
	// @TODO: need to implement
	return nil
}

func (s *structure) Update(key int, value int) error {
	// @TODO: need to implement
	return nil
}

func (s *structure) Get(key string) (int, error) {
	// @TODO: need to implement
	return 0, nil
}

func New(keyValues map[string]int) *structure {
	s := &structure{}

	// @TODO: need to implement

	return s
}
