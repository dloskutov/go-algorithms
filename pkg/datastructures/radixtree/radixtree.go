package radixtree

type RadixTree struct{}

func (t *RadixTree) Insert(key string) error {
	// @TODO: need to implement
	return nil
}

func (t *RadixTree) Remove(key string) error {
	// @TODO: need to implement
	return nil
}

func (t *RadixTree) Contains(key string) bool {
	// @TODO: need to implement
	return false
}

func (t *RadixTree) FindLongestPrefix(key string) string {
	// @TODO: need to implement
	return ""
}

func (t *RadixTree) KeysStartingWith(prefix string) []string {
	// @TODO: need to implement
	return []string{}
}

func New(data []string) (*RadixTree, error) {
	t := &RadixTree{}
	for i := range data {
		err := t.Insert(data[i])
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
