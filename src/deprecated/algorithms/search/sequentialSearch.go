package search

import "fmt"

type node struct {
	key   string
	value int64
	node  *node
}

// SequentialSearch - sequential search by key
type SequentialSearch struct {
	first *node
}

// Put - put value by key
func (ss *SequentialSearch) Put(key string, value int64) {
	currentNode := ss.first
	for currentNode != nil {
		if currentNode.key == key {
			currentNode.value = value
			return
		}
		currentNode = currentNode.node
	}
	ss.first = &node{
		key:   key,
		value: value,
		node:  ss.first,
	}
}

// Get - search value by key
func (ss *SequentialSearch) Get(key string) (int64, error) {
	currentNode := ss.first
	for currentNode != nil {
		if currentNode.key == key {
			return currentNode.value, nil
		}
		currentNode = currentNode.node
	}
	return 0, fmt.Errorf("there is no value associated with key: %s", key)
}
