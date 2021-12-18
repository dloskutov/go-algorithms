package binarysearchtree

import "fmt"

var ErrInvalidKey = fmt.Errorf("invalid key")

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
	// @TODO: need to implement
	return nil
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
