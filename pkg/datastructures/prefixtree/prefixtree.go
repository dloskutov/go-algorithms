package prefixtree

type node struct {
	symbol   rune
	value    interface{}
	children map[rune]*node
}

type PrefixTree struct {
	root *node
}

func (t *PrefixTree) Insert(key string, value interface{}) error {
	// @TODO: need to implement
	return nil
}

func (t *PrefixTree) Get(key string) (interface{}, error) {
	// @TODO: need to implement
	return nil, nil
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

func New(data map[string]interface{}) (*PrefixTree, error) {
	t := &PrefixTree{
		root: &node{
			symbol:   rune(0),
			value:    nil,
			children: make(map[rune]*node),
		},
	}
	for key, value := range data {
		err := t.Insert(key, value)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
