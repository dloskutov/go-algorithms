package prefixtree

import "fmt"

var ErrInvalidKey = fmt.Errorf("invlid key")

type nodeWithPrefix struct {
	node   *node
	prefix string
}

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

// @TODO: is it the right implementation?
func (t *PrefixTree) FindLongestPrefix(key string) string {
	result := t.KeysStartingWith(key)

	if len(result) == 0 {
		return ""
	}

	index := 0
	for i := range result {
		if len(result[i]) > len(result[index]) {
			index = i
		}
	}

	return result[index]
}

func (t *PrefixTree) KeysStartingWith(prefix string) []string {
	if prefix == "" {
		return []string{}
	}

	index := 0
	length := len(prefix)

	parentNode := t.root
	for index < length {
		char := prefix[index]
		if parentNode.children[char] != nil {
			parentNode = parentNode.children[char]
		} else {
			return []string{}
		}
		index++
	}

	result := make([]string, 0)
	nodesWithPrefix := []*nodeWithPrefix{{
		node:   parentNode,
		prefix: prefix[:len(prefix)-1],
	}}
	for len(nodesWithPrefix) != 0 {
		lastIndex := len(nodesWithPrefix) - 1
		n := nodesWithPrefix[lastIndex]

		if n.node.value != nil {
			result = append(result, n.prefix+string(n.node.symbol))
		}

		nodesWithPrefix = nodesWithPrefix[:lastIndex]
		for i := range n.node.children {
			nodesWithPrefix = append(nodesWithPrefix, &nodeWithPrefix{
				node:   n.node.children[i],
				prefix: n.prefix + string(n.node.symbol),
			})
		}
	}

	return result
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
