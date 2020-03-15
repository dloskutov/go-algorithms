package search

import (
	"fmt"
	"strings"
)

type nodeBST struct {
	key        string
	value      int64
	leftChild  *nodeBST
	rightChild *nodeBST
}

// BinarySearchTree - sequential search by key
type BinarySearchTree struct {
	root *nodeBST
}

// Put - put value by key to the tree
func (bs *BinarySearchTree) Put(key string, value int64) {
	if bs.root == nil {
		bs.root = &nodeBST{
			key:        key,
			value:      value,
			leftChild:  nil,
			rightChild: nil,
		}
		return
	}

	node := bs.root

	for {
		compareKey := strings.Compare(key, node.key)

		if compareKey == 0 {
			node.value = value
			break
		}
		if compareKey < 0 {
			if node.leftChild == nil {
				node.leftChild = &nodeBST{
					key:        key,
					value:      value,
					leftChild:  nil,
					rightChild: nil,
				}
				break
			}
			node = node.leftChild
			continue
		}
		if compareKey > 0 {
			if node.rightChild == nil {
				node.rightChild = &nodeBST{
					key:        key,
					value:      value,
					leftChild:  nil,
					rightChild: nil,
				}
				break
			}
			node = node.rightChild
			continue
		}
	}
}

// Get - search value by key
func (bs *BinarySearchTree) Get(key string) (int64, error) {
	node := bs.root

	for {
		if node == nil {
			break
		}

		compareKey := strings.Compare(key, node.key)
		if compareKey == 0 {
			return node.value, nil
		}
		if compareKey < 0 {
			node = node.leftChild
			continue
		}
		if compareKey > 0 {
			node = node.rightChild
			continue
		}
	}

	return 0, fmt.Errorf("there is no value associated with key: %s", key)
}
