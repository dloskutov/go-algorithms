package search

import (
	"fmt"
	"strings"
)

const nodeRed = 1
const nodeBlack = 2

type redBlackNode struct {
	color      int
	key        string
	value      int64
	childLeft  *redBlackNode
	childRight *redBlackNode
	parent     *redBlackNode
}

// RedBlackBinarySearchTree - binary search tree with red-black tree balancing
type RedBlackBinarySearchTree struct {
	root *redBlackNode
}

// Put - RedBlackBinarySearchTree put value by key
func (tree *RedBlackBinarySearchTree) Put(key string, value int64) {
	node := tree.root

	if node == nil {
		tree.root = &redBlackNode{
			color:      nodeBlack,
			key:        key,
			value:      value,
			childLeft:  nil,
			childRight: nil,
			parent:     nil,
		}
		return
	}

	for node != nil {
		compareKey := strings.Compare(key, node.key)

		if compareKey == 0 {
			node.value = value
			break
		}

		if compareKey < 0 {
			if node.childLeft != nil {
				node = node.childLeft
			} else {
				node.childLeft = &redBlackNode{
					color:      nodeRed,
					key:        key,
					value:      value,
					childLeft:  nil,
					childRight: nil,
					parent:     node,
				}
				tree.rebalance(node.childLeft)
				break
			}
		} else {
			if node.childRight != nil {
				node = node.childRight
			} else {
				node.childRight = &redBlackNode{
					color:      nodeRed,
					key:        key,
					value:      value,
					childLeft:  nil,
					childRight: nil,
					parent:     node,
				}
				tree.rebalance(node.childRight)
				break
			}
		}
	}
}

// Get - search value by key
func (tree *RedBlackBinarySearchTree) Get(key string) (int64, error) {
	node := tree.root

	for node != nil {
		if node.key == key {
			return node.value, nil
		}
		if strings.Compare(key, node.key) < 0 {
			node = node.childLeft
		} else {
			node = node.childRight
		}
	}

	return 0, fmt.Errorf("there is no value associated with key: %s", key)
}

func (tree *RedBlackBinarySearchTree) rebalance(node *redBlackNode) {
	parent := node.parent
	grandParent := getGrandParentNode(node)
	uncle := getUncleNode(node)

	if parent == nil {
		node.color = nodeBlack
		return
	}

	if parent.color == nodeBlack || grandParent == nil {
		parent.color = nodeBlack
		return
	}

	if isNodeRightChildOfParent(node) && isNodeLeftChildOfParent(node.parent) {
		tree.rotateLeft(node)
		tree.rebalance(parent)
		return
	}

	if isNodeLeftChildOfParent(node) && isNodeRightChildOfParent(node.parent) {
		tree.rotateRight(node)
		tree.rebalance(parent)
		return
	}

	if uncle == nil {
		if isNodeLeftChildOfParent(parent) {
			tree.rotateRight(parent)
			parent.color = nodeBlack
			grandParent.color = nodeRed
			tree.rebalance(parent)
			return
		}
		if isNodeRightChildOfParent(parent) {
			tree.rotateLeft(parent)
			parent.color = nodeBlack
			grandParent.color = nodeRed
			tree.rebalance(parent)
			return
		}
	}
	if uncle.color == nodeBlack {
		if isNodeLeftChildOfParent(parent) {
			tree.rotateRight(parent)
			parent.color = nodeBlack
			grandParent.color = nodeRed
			return
		}
		if isNodeRightChildOfParent(parent) {
			tree.rotateLeft(parent)
			parent.color = nodeBlack
			grandParent.color = nodeRed
			return
		}
	}
	if uncle.color == nodeRed {
		parent.color = nodeBlack
		uncle.color = nodeBlack
		grandParent.color = nodeRed
		tree.rebalance(grandParent)
		return
	}
}

func getGrandParentNode(node *redBlackNode) *redBlackNode {
	if node == nil || node.parent == nil {
		return nil
	}

	return node.parent.parent
}

func getUncleNode(node *redBlackNode) *redBlackNode {
	grandParent := getGrandParentNode(node)
	if grandParent == nil {
		return nil
	}
	if grandParent.childLeft == node.parent {
		return grandParent.childRight
	}
	if grandParent.childRight == node.parent {
		return grandParent.childLeft
	}
	return nil
}

func isNodeRightChildOfParent(node *redBlackNode) bool {
	if node != nil && node.parent != nil && node.parent.childRight == node {
		return true
	}

	return false
}

func isNodeLeftChildOfParent(node *redBlackNode) bool {
	if node != nil && node.parent != nil && node.parent.childLeft == node {
		return true
	}

	return false
}

// make right child = parent, parent = left child
func (tree *RedBlackBinarySearchTree) rotateLeft(node *redBlackNode) {
	oldRightChild := node
	oldParent := node.parent
	oldGrandParent := oldParent.parent

	isLeftChild := isNodeLeftChildOfParent(oldParent)
	isRightChild := isNodeRightChildOfParent(oldParent)

	oldRightChild.parent = oldParent.parent
	oldParent.parent = oldRightChild

	oldParent.childRight = oldRightChild.childLeft
	oldRightChild.childLeft = oldParent

	if oldParent.childRight != nil {
		oldParent.childRight.parent = oldParent
	}

	if oldGrandParent != nil {
		if isLeftChild {
			oldGrandParent.childLeft = oldRightChild
			return
		}
		if isRightChild {
			oldGrandParent.childRight = oldRightChild
			return
		}
	} else {
		tree.root = oldRightChild
	}
}

// make left child = parent, parent = right child
func (tree *RedBlackBinarySearchTree) rotateRight(node *redBlackNode) {
	oldLeftChild := node
	oldParent := node.parent
	oldGrandParent := oldParent.parent

	isLeftChild := isNodeLeftChildOfParent(oldParent)
	isRightChild := isNodeRightChildOfParent(oldParent)

	oldLeftChild.parent = oldParent.parent
	oldParent.parent = oldLeftChild

	oldParent.childLeft = oldLeftChild.childRight
	oldLeftChild.childRight = oldParent

	if oldParent.childLeft != nil {
		oldParent.childLeft.parent = oldParent
	}

	if oldGrandParent != nil {
		if isLeftChild {
			oldGrandParent.childLeft = oldLeftChild
			return
		}
		if isRightChild {
			oldGrandParent.childRight = oldLeftChild
			return
		}
	} else {
		tree.root = oldLeftChild
	}
}
