package search

import (
	"fmt"
	"strings"
)

type nodeFour struct {
	leftKey          string
	middleKey        string
	rightKey         string
	leftValue        int64
	middleValue      int64
	rightValue       int64
	leftChild        interface{}
	leftMiddleChild  interface{}
	rightMiddleChild interface{}
	rightChild       interface{}
	parent           interface{}
}

type nodeThree struct {
	leftKey     string
	rightKey    string
	leftValue   int64
	rightValue  int64
	leftChild   interface{}
	middleChild interface{}
	rightChild  interface{}
	parent      interface{}
}

type nodeTwo struct {
	key        string
	value      int64
	leftChild  interface{}
	rightChild interface{}
	parent     interface{}
}

// TwoThreeSearchTree - two-three search tree
type TwoThreeSearchTree struct {
	root interface{}
}

// Put - put value by key to the tree
func (tree *TwoThreeSearchTree) Put(key string, value int64) {
	nodeRaw := tree.root

	if nodeRaw == nil {
		tree.root = &nodeTwo{
			key:        key,
			value:      value,
			leftChild:  nil,
			rightChild: nil,
			parent:     nil,
		}
		return
	}

	for {
		switch node := nodeRaw.(type) {
		case *nodeTwo:
			compareKey := strings.Compare(key, node.key)

			if compareKey == 0 {
				node.value = value
				return
			}
			if compareKey < 0 && node.leftChild != nil {
				nodeRaw = node.leftChild
				continue
			}
			if compareKey > 0 && node.rightChild != nil {
				nodeRaw = node.rightChild
				continue
			}

			var newNode *nodeThree
			if compareKey < 0 {
				newNode = &nodeThree{
					leftKey:     key,
					rightKey:    node.key,
					leftValue:   value,
					rightValue:  node.value,
					leftChild:   nil,
					middleChild: node.leftChild,
					rightChild:  node.rightChild,
					parent:      node.parent,
				}
			} else {
				newNode = &nodeThree{
					leftKey:     node.key,
					rightKey:    key,
					leftValue:   node.value,
					rightValue:  value,
					leftChild:   node.leftChild,
					middleChild: node.rightChild,
					rightChild:  nil,
					parent:      node.parent,
				}
			}

			updateParent(newNode.leftChild, newNode)
			updateParent(newNode.middleChild, newNode)
			updateParent(newNode.rightChild, newNode)

			if node.parent == nil {
				tree.root = newNode
			}
			updateParentChildNode(node, newNode)
			return
		case *nodeThree:
			leftCompareKey := strings.Compare(key, node.leftKey)
			rightCompareKey := strings.Compare(key, node.rightKey)

			if leftCompareKey == 0 {
				node.leftValue = value
				return
			}
			if rightCompareKey == 0 {
				node.rightValue = value
				return
			}
			if leftCompareKey < 0 && node.leftChild != nil {
				nodeRaw = node.leftChild
				continue
			}
			if rightCompareKey > 0 && node.rightChild != nil {
				nodeRaw = node.rightChild
				continue
			}
			if node.middleChild != nil {
				nodeRaw = node.middleChild
				continue
			}

			var newNode *nodeFour
			if leftCompareKey < 0 {
				newNode = &nodeFour{
					leftKey:          key,
					middleKey:        node.leftKey,
					rightKey:         node.rightKey,
					leftValue:        value,
					middleValue:      node.leftValue,
					rightValue:       node.rightValue,
					leftChild:        nil,
					leftMiddleChild:  nil,
					rightMiddleChild: node.middleChild,
					rightChild:       node.rightChild,
					parent:           node.parent,
				}
			} else if rightCompareKey > 0 {
				newNode = &nodeFour{
					leftKey:          node.leftKey,
					middleKey:        node.rightKey,
					rightKey:         key,
					leftValue:        node.leftValue,
					middleValue:      node.rightValue,
					rightValue:       value,
					leftChild:        node.leftChild,
					leftMiddleChild:  node.middleChild,
					rightMiddleChild: nil,
					rightChild:       nil,
					parent:           node.parent,
				}
			} else {
				newNode = &nodeFour{
					leftKey:          node.leftKey,
					middleKey:        key,
					rightKey:         node.rightKey,
					leftValue:        node.leftValue,
					middleValue:      value,
					rightValue:       node.rightValue,
					leftChild:        node.leftChild,
					leftMiddleChild:  nil,
					rightMiddleChild: nil,
					rightChild:       node.rightChild,
					parent:           node.parent,
				}
			}

			updateParent(newNode.leftChild, newNode)
			updateParent(newNode.leftMiddleChild, newNode)
			updateParent(newNode.rightMiddleChild, newNode)
			updateParent(newNode.rightChild, newNode)

			if node.parent == nil {
				tree.root = newNode
			}

			updateParentChildNode(node, newNode)
			tree.rebalanceTree(newNode)
			return
		default:
			// skip
		}
	}
}

// Get - search value by key
func (tree *TwoThreeSearchTree) Get(key string) (int64, error) {
	nodeRaw := tree.root

loop:
	for {
		if nodeRaw == nil {
			break loop
		}
		switch node := nodeRaw.(type) {
		case *nodeTwo:
			compareKey := strings.Compare(key, node.key)

			if compareKey == 0 {
				return node.value, nil
			}
			if compareKey < 0 {
				nodeRaw = node.leftChild
				continue
			}
			if compareKey > 0 {
				nodeRaw = node.rightChild
				continue
			}
		case *nodeThree:
			leftCompareKey := strings.Compare(key, node.leftKey)
			rightCompareKey := strings.Compare(key, node.rightKey)

			if leftCompareKey == 0 {
				return node.leftValue, nil
			}
			if rightCompareKey == 0 {
				return node.rightValue, nil
			}

			if leftCompareKey < 0 {
				nodeRaw = node.leftChild
				continue
			}
			if rightCompareKey > 0 {
				nodeRaw = node.rightChild
				continue
			}
			nodeRaw = node.middleChild
		default:
			return 0, fmt.Errorf("unable to convert raw node to node: %v", nodeRaw)
		}
	}

	return 0, fmt.Errorf("there is no value associated with key: %s", key)
}

func (tree *TwoThreeSearchTree) rebalanceTree(node *nodeFour) {
	nodeParentRaw := node.parent

	if nodeParentRaw == nil {
		rootNode := &nodeTwo{
			key:        node.middleKey,
			value:      node.middleValue,
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
		}
		leftChild := &nodeTwo{
			key:        node.leftKey,
			value:      node.leftValue,
			parent:     rootNode,
			leftChild:  node.leftChild,
			rightChild: node.leftMiddleChild,
		}
		rightChild := &nodeTwo{
			key:        node.rightKey,
			value:      node.rightValue,
			parent:     rootNode,
			leftChild:  node.rightMiddleChild,
			rightChild: node.rightChild,
		}

		updateParent(leftChild.leftChild, leftChild)
		updateParent(leftChild.rightChild, leftChild)
		updateParent(rightChild.leftChild, rightChild)
		updateParent(rightChild.rightChild, rightChild)

		rootNode.leftChild = leftChild
		rootNode.rightChild = rightChild
		tree.root = rootNode
		return
	}

	switch parentNode := nodeParentRaw.(type) {
	case *nodeTwo:
		// convert nodeTwo parent to nodeThree parent
		compareKey := strings.Compare(node.middleKey, parentNode.key)

		var newParentNode *nodeThree
		if compareKey < 0 {
			newParentNode = &nodeThree{
				leftKey:     node.middleKey,
				rightKey:    parentNode.key,
				leftValue:   node.middleValue,
				rightValue:  parentNode.value,
				leftChild:   nil,
				middleChild: nil,
				rightChild:  nil,
				parent:      parentNode.parent,
			}
			leftChild := &nodeTwo{
				key:        node.leftKey,
				value:      node.leftValue,
				parent:     newParentNode,
				leftChild:  node.leftChild,
				rightChild: node.leftMiddleChild,
			}
			middleChild := &nodeTwo{
				key:        node.rightKey,
				value:      node.rightValue,
				parent:     newParentNode,
				leftChild:  node.rightMiddleChild,
				rightChild: node.rightChild,
			}

			newParentNode.leftChild = leftChild
			newParentNode.middleChild = middleChild
			newParentNode.rightChild = parentNode.rightChild

			updateParent(newParentNode.rightChild, newParentNode)
			updateParent(leftChild.leftChild, leftChild)
			updateParent(leftChild.rightChild, leftChild)
			updateParent(middleChild.leftChild, middleChild)
			updateParent(middleChild.rightChild, middleChild)
		} else {
			newParentNode = &nodeThree{
				leftKey:     parentNode.key,
				rightKey:    node.middleKey,
				leftValue:   parentNode.value,
				rightValue:  node.middleValue,
				leftChild:   nil,
				middleChild: nil,
				rightChild:  nil,
				parent:      parentNode.parent,
			}
			middleChild := &nodeTwo{
				key:        node.leftKey,
				value:      node.leftValue,
				parent:     newParentNode,
				leftChild:  node.leftChild,
				rightChild: node.leftMiddleChild,
			}
			rightChild := &nodeTwo{
				key:        node.rightKey,
				value:      node.rightValue,
				parent:     newParentNode,
				leftChild:  node.rightMiddleChild,
				rightChild: node.rightChild,
			}

			newParentNode.leftChild = parentNode.leftChild
			newParentNode.middleChild = middleChild
			newParentNode.rightChild = rightChild

			updateParent(newParentNode.leftChild, newParentNode)
			updateParent(middleChild.leftChild, middleChild)
			updateParent(middleChild.rightChild, middleChild)
			updateParent(rightChild.leftChild, rightChild)
			updateParent(rightChild.rightChild, rightChild)
		}

		if newParentNode.parent == nil {
			tree.root = newParentNode
		}

		updateParentChildNode(parentNode, newParentNode)
	case *nodeThree:
		leftCompareKey := strings.Compare(node.middleKey, parentNode.leftKey)
		rightCompareKey := strings.Compare(node.middleKey, parentNode.rightKey)

		var newParentNode *nodeFour
		if leftCompareKey < 0 {
			newParentNode = &nodeFour{
				leftKey:          node.middleKey,
				middleKey:        parentNode.leftKey,
				rightKey:         parentNode.rightKey,
				leftValue:        node.middleValue,
				middleValue:      parentNode.leftValue,
				rightValue:       parentNode.rightValue,
				leftChild:        nil,
				leftMiddleChild:  nil,
				rightMiddleChild: parentNode.middleChild,
				rightChild:       parentNode.rightChild,
				parent:           parentNode.parent,
			}
			leftChild := &nodeTwo{
				key:        node.leftKey,
				value:      node.leftValue,
				parent:     newParentNode,
				leftChild:  node.leftChild,
				rightChild: node.leftMiddleChild,
			}
			leftMiddleChild := &nodeTwo{
				key:        node.rightKey,
				value:      node.rightValue,
				parent:     newParentNode,
				leftChild:  node.rightMiddleChild,
				rightChild: node.rightChild,
			}

			newParentNode.leftChild = leftChild
			newParentNode.leftMiddleChild = leftMiddleChild

			updateParent(parentNode.middleChild, newParentNode)
			updateParent(parentNode.rightChild, newParentNode)
			updateParent(leftChild.leftChild, leftChild)
			updateParent(leftChild.rightChild, leftChild)
			updateParent(leftMiddleChild.leftChild, leftMiddleChild)
			updateParent(leftMiddleChild.rightChild, leftMiddleChild)
		} else if rightCompareKey > 0 {
			newParentNode = &nodeFour{
				leftKey:          parentNode.leftKey,
				middleKey:        parentNode.rightKey,
				rightKey:         node.middleKey,
				leftValue:        parentNode.leftValue,
				middleValue:      parentNode.rightValue,
				rightValue:       node.middleValue,
				leftChild:        parentNode.leftChild,
				leftMiddleChild:  parentNode.middleChild,
				rightMiddleChild: nil,
				rightChild:       nil,
				parent:           parentNode.parent,
			}
			rightMiddleChild := &nodeTwo{
				key:        node.leftKey,
				value:      node.leftValue,
				parent:     newParentNode,
				leftChild:  node.leftChild,
				rightChild: node.leftMiddleChild,
			}
			rightChild := &nodeTwo{
				key:        node.rightKey,
				value:      node.rightValue,
				parent:     newParentNode,
				leftChild:  node.rightMiddleChild,
				rightChild: node.rightChild,
			}

			newParentNode.rightMiddleChild = rightMiddleChild
			newParentNode.rightChild = rightChild

			updateParent(parentNode.leftChild, newParentNode)
			updateParent(parentNode.middleChild, newParentNode)
			updateParent(rightMiddleChild.leftChild, rightMiddleChild)
			updateParent(rightMiddleChild.rightChild, rightMiddleChild)
			updateParent(rightChild.leftChild, rightChild)
			updateParent(rightChild.rightChild, rightChild)
		} else {
			newParentNode = &nodeFour{
				leftKey:          parentNode.leftKey,
				middleKey:        node.middleKey,
				rightKey:         parentNode.rightKey,
				leftValue:        parentNode.leftValue,
				middleValue:      node.middleValue,
				rightValue:       parentNode.rightValue,
				leftChild:        parentNode.leftChild,
				leftMiddleChild:  nil,
				rightMiddleChild: nil,
				rightChild:       parentNode.rightChild,
				parent:           parentNode.parent,
			}
			leftMiddleChild := &nodeTwo{
				key:        node.leftKey,
				value:      node.leftValue,
				parent:     newParentNode,
				leftChild:  node.leftChild,
				rightChild: node.leftMiddleChild,
			}
			rightMiddleChild := &nodeTwo{
				key:        node.rightKey,
				value:      node.rightValue,
				parent:     newParentNode,
				leftChild:  node.rightMiddleChild,
				rightChild: node.rightChild,
			}

			newParentNode.leftMiddleChild = leftMiddleChild
			newParentNode.rightMiddleChild = rightMiddleChild

			updateParent(parentNode.leftChild, newParentNode)
			updateParent(parentNode.rightChild, newParentNode)
			updateParent(leftMiddleChild.leftChild, leftMiddleChild)
			updateParent(leftMiddleChild.rightChild, leftMiddleChild)
			updateParent(rightMiddleChild.leftChild, rightMiddleChild)
			updateParent(rightMiddleChild.rightChild, rightMiddleChild)
		}

		if newParentNode.parent == nil {
			tree.root = newParentNode
		}

		updateParentChildNode(parentNode, newParentNode)
		tree.rebalanceTree(newParentNode)
	}
}

func updateParent(currentNodeRaw interface{}, parent interface{}) {
	if currentNodeRaw != nil {
		switch currentNode := currentNodeRaw.(type) {
		case *nodeTwo:
			currentNode.parent = parent
		case *nodeThree:
			currentNode.parent = parent
		}
	}
}

func updateParentChildNode(currentNodeRaw interface{}, newNode interface{}) {
	if currentNodeRaw != nil {
		switch currentNode := currentNodeRaw.(type) {
		case *nodeTwo:
			if currentNode.parent != nil {
				switch parentNode := currentNode.parent.(type) {
				case *nodeTwo:
					if parentNode.leftChild == currentNodeRaw {
						parentNode.leftChild = newNode
					} else if parentNode.rightChild == currentNodeRaw {
						parentNode.rightChild = newNode
					}
				case *nodeThree:
					if parentNode.leftChild == currentNodeRaw {
						parentNode.leftChild = newNode
					} else if parentNode.middleChild == currentNodeRaw {
						parentNode.middleChild = newNode
					} else if parentNode.rightChild == currentNodeRaw {
						parentNode.rightChild = newNode
					}
				}
			}
		case *nodeThree:
			if currentNode.parent != nil {
				switch parentNode := currentNode.parent.(type) {
				case *nodeTwo:
					if parentNode.leftChild == currentNodeRaw {
						parentNode.leftChild = newNode
					}
					if parentNode.rightChild == currentNodeRaw {
						parentNode.rightChild = newNode
					}
				case *nodeThree:
					if parentNode.leftChild == currentNodeRaw {
						parentNode.leftChild = newNode
					}
					if parentNode.middleChild == currentNodeRaw {
						parentNode.middleChild = newNode
					}
					if parentNode.rightChild == currentNodeRaw {
						parentNode.rightChild = newNode
					}
				}
			}
		}
	}
}
