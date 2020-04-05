package search

import "fmt"

// prime number (better use 997)
const lengthOfArray = 3

type separateChainingNode struct {
	key      string
	value    int64
	nextNode *separateChainingNode
}

// SeparateChainingHashTable - search key in separate chaining hash table
type SeparateChainingHashTable struct {
	array []*separateChainingNode
}

// Put - put value by key
func (hashTable *SeparateChainingHashTable) Put(key string, value int64) {
	if hashTable.array == nil {
		hashTable.array = make([]*separateChainingNode, lengthOfArray)
	}
	hashIndex := hashFunc(key)
	list := hashTable.array[hashIndex]
	if list == nil {
		hashTable.array[hashIndex] = &separateChainingNode{
			key:      key,
			value:    value,
			nextNode: nil,
		}
		return
	}
	node := list
	for node != nil {
		if node.key == key {
			node.value = value
			return
		}
		if node.nextNode == nil {
			node.nextNode = &separateChainingNode{
				key:      key,
				value:    value,
				nextNode: nil,
			}
			return
		}
		node = node.nextNode
	}
}

// Get - search value by key
func (hashTable *SeparateChainingHashTable) Get(key string) (int64, error) {
	if hashTable.array == nil {
		return 0, fmt.Errorf("there is no value for key: %s", key)
	}
	hashIndex := hashFunc(key)
	list := hashTable.array[hashIndex]
	if list == nil {
		return 0, fmt.Errorf("there is no value for key: %s", key)
	}
	node := list
	for node != nil {
		if node.key == key {
			return node.value, nil
		}
		node = node.nextNode
	}

	return 0, fmt.Errorf("there is no value for key: %s", key)
}

func hashFunc(key string) uint64 {
	var value uint64
	for _, letter := range key {
		value += uint64(letter)
	}
	return value % lengthOfArray
}
