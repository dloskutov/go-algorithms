package binarysearchtree

import "fmt"

var (
	ErrInvalidKey   = fmt.Errorf("invalid key")
	ErrNodeNotFound = fmt.Errorf("node not found")
)

type node struct {
	key       int
	value     interface{}
	leftNode  *node
	rightNode *node
}

type BinarySearchTree struct {
	root *node
}

func (bst *BinarySearchTree) Put(key int, value interface{}) error {
	if bst.root == nil {
		bst.root = &node{
			key:   key,
			value: value,
		}
		return nil
	}

	currentNode := bst.root
	parentNode := currentNode
	for currentNode != nil {
		parentNode = currentNode
		if key > currentNode.key {
			currentNode = currentNode.rightNode
		} else {
			currentNode = currentNode.leftNode
		}
	}
	if key > parentNode.key {
		parentNode.rightNode = &node{
			key:   key,
			value: value,
		}
	} else {
		parentNode.leftNode = &node{
			key:   key,
			value: value,
		}
	}

	return nil
}

func (bst *BinarySearchTree) Search(key int) (interface{}, error) {
	parentNode := bst.root
	for parentNode != nil {
		if parentNode.key == key {
			return parentNode.value, nil
		} else if key > parentNode.key {
			parentNode = parentNode.rightNode
		} else {
			parentNode = parentNode.leftNode
		}
	}

	return nil, ErrInvalidKey
}

func (bst *BinarySearchTree) Remove(key int) error {
	parentNode := bst.root
	currentNode := parentNode

	for currentNode != nil {
		if currentNode.key == key {
			// case 1
			if currentNode == bst.root {
				// @TODO: need to implement
				return nil
			}

			// case 2 (both child nodes are empty)
			if currentNode.leftNode == nil && currentNode.rightNode == nil {
				return removeChild(parentNode, currentNode)
			}

			// case 3 (leftNode is not exist)
			if currentNode.leftNode == nil && currentNode.rightNode != nil {
				return replaceChild(parentNode, currentNode, currentNode.rightNode)
			}

			// case 4 (rightNode is not exist)
			if currentNode.rightNode == nil && currentNode.leftNode != nil {
				return replaceChild(parentNode, currentNode, currentNode.leftNode)
			}

			// @TODO: need to implement
			// other cases

			return nil
		}

		parentNode = currentNode
		if key > currentNode.key {
			currentNode = currentNode.rightNode
		} else {
			currentNode = currentNode.leftNode
		}
	}

	return ErrInvalidKey
}

func New(data map[int]interface{}) (*BinarySearchTree, error) {
	bst := &BinarySearchTree{}
	for key, value := range data {
		err := bst.Put(key, value)
		if err != nil {
			return nil, err
		}
	}
	return bst, nil
}

func removeChild(parentNode *node, childNode *node) error {
	if parentNode == nil {
		return ErrNodeNotFound
	}
	if parentNode.leftNode == childNode {
		parentNode.leftNode = nil
		return nil
	}
	if parentNode.rightNode == childNode {
		parentNode.rightNode = nil
		return nil
	}
	return ErrNodeNotFound
}

func replaceChild(parentNode *node, oldChildNode *node, newChildNode *node) error {
	if parentNode == nil {
		return ErrNodeNotFound
	}
	if parentNode.leftNode == oldChildNode {
		parentNode.leftNode = newChildNode
		return nil
	}
	if parentNode.rightNode == oldChildNode {
		parentNode.rightNode = newChildNode
		return nil
	}

	return ErrNodeNotFound
}
