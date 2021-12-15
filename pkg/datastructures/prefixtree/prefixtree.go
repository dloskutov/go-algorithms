package prefixtree

type PrefixTree struct{}

func (t *PrefixTree) Insert(key string) error {
	// @TODO: need to implement
	return nil
}

func (t *PrefixTree) Remove(key string) error {
	// @TODO: need to implement
	return nil
}

func (t *PrefixTree) Contains(key string) bool {
	// @TODO: need to implement
	return false
}

func (t *PrefixTree) FindLongestPrefix(key string) string {
	// @TODO: need to implement
	return ""
}

func (t *PrefixTree) KeysStartingWith(prefix string) []string {
	// @TODO: need to implement
	return []string{}
}

func New(data []string) (*PrefixTree, error) {
	t := &PrefixTree{}
	for i := range data {
		err := t.Insert(data[i])
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
