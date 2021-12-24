package prefixtree

import "fmt"

var ErrInvalidKey = fmt.Errorf("invlid key")

type node struct {
	symbol   byte
	value    interface{}
	children map[byte]*node
}

type PrefixTree struct {
	root *node
}

func (t *PrefixTree) Insert(key string, value interface{}) error {
	index := 0
	length := len(key)

	parentNode := t.root
	for index < length {
		char := key[index]
		if parentNode.children[char] == nil {
			parentNode.children[char] = &node{
				symbol:   char,
				value:    nil,
				children: make(map[byte]*node),
			}
		}
		parentNode = parentNode.children[char]
		index++
	}
	parentNode.value = value

	return nil
}

func (t *PrefixTree) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}

	index := 0
	length := len(key)

	parentNode := t.root
	for index < length {
		char := key[index]
		if parentNode.children[char] != nil {
			parentNode = parentNode.children[char]
		} else {
			return nil, ErrInvalidKey
		}
		index++
	}

	if parentNode.value != nil {
		return parentNode.value, nil
	}

	return nil, ErrInvalidKey
}

func (t *PrefixTree) Remove(key string) error {
	if key == "" {
		return ErrInvalidKey
	}

	index := 0
	length := len(key)

	parentNode := t.root
	for index < length {
		char := key[index]
		if parentNode.children[char] != nil {
			parentNode = parentNode.children[char]
		} else {
			return ErrInvalidKey
		}
		index++
	}

	// @TODO: need delete unused childred too!!
	if parentNode.value != nil {
		parentNode.value = nil
		return nil
	}

	return ErrInvalidKey
}

func (t *PrefixTree) Contains(key string) bool {
	value, err := t.Get(key)
	return err == nil && value != nil
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
			symbol:   0,
			value:    nil,
			children: make(map[byte]*node),
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
