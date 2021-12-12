package binaryheap

import (
	"fmt"

	"github.com/dloskutov/go-algorithms/src/datastructure/array"
)

var ErrBinaryHeapIsEmpty = fmt.Errorf("binary heap is empty")
var ErrNotFound = fmt.Errorf("value with priority not found")

type node struct {
	priority int
	value    interface{}
}

type heapType int

const (
	Min heapType = 1
	Max heapType = 2
)

type BinaryHeap struct {
	heapType  heapType
	direction int
	array     *array.Array
}

func (s *BinaryHeap) Insert(priority int, value interface{}) error {
	s.array.Add(&node{
		priority: priority,
		value:    value,
	})

	index := s.array.Size() - 1
	return s.bubbleUp(index)
}

func (h *BinaryHeap) compare(leftChildIndex int, rightChildIndex int) int {
	leftChildRaw, err := h.array.Get(leftChildIndex)
	if err != nil {
		return 1 // go to right child node
	}
	leftChild := leftChildRaw.(*node)
	rightChildRaw, err := h.array.Get(rightChildIndex)
	if err != nil {
		return -1 // go to left child node
	}
	rightChild := rightChildRaw.(*node)
	if h.direction*rightChild.priority > h.direction*leftChild.priority {
		return 1
	} else {
		return -1
	}
}

func (h *BinaryHeap) bubbleDown(index int) error {
	size := h.array.Size()
	parentIndex := index
	leftChildIndex := (parentIndex * 2) + 1
	rightChildIndex := (parentIndex * 2) + 2

	for {
		parentNodeRaw, err := h.array.Get(parentIndex)
		if err != nil {
			return err
		}
		parentNode := parentNodeRaw.(*node)
		compare := h.compare(leftChildIndex, rightChildIndex)
		if compare > 0 {
			goto right
		} else if compare < 0 {
			goto left
		}

	left:
		if leftChildIndex < size {
			leftChildNodeRaw, err := h.array.Get(leftChildIndex)
			if err != nil {
				return err
			}
			leftChildNode := leftChildNodeRaw.(*node)
			if (h.direction * leftChildNode.priority) > (h.direction * parentNode.priority) {
				err = h.array.Update(leftChildIndex, parentNode)
				if err != nil {
					return err
				}
				err = h.array.Update(parentIndex, leftChildNode)
				if err != nil {
					return err
				}

				parentIndex = leftChildIndex
				leftChildIndex = (parentIndex * 2) + 1
				rightChildIndex = (parentIndex * 2) + 2
				continue
			}
		}

	right:
		if rightChildIndex < size {
			rightChildNodeRaw, err := h.array.Get(rightChildIndex)
			if err != nil {
				return err
			}
			rightChildNode := rightChildNodeRaw.(*node)
			if (h.direction * rightChildNode.priority) > (h.direction * parentNode.priority) {
				err = h.array.Update(rightChildIndex, parentNode)
				if err != nil {
					return err
				}
				err = h.array.Update(parentIndex, rightChildNode)
				if err != nil {
					return err
				}

				parentIndex = rightChildIndex
				leftChildIndex = (parentIndex * 2) + 1
				rightChildIndex = (parentIndex * 2) + 2
				continue
			}
		}
		break
	}

	return nil
}

func (h *BinaryHeap) bubbleUp(index int) error {
	currentNodeRaw, err := h.array.Get(index)
	if err != nil {
		return err
	}
	currentNode := currentNodeRaw.(*node)

	currentIndex := index
	parentIndex := getParentIndex(currentIndex)

	for parentIndex >= 0 {
		parentNodeRaw, err := h.array.Get(parentIndex)
		if err != nil {
			return err
		}
		parentNode := parentNodeRaw.(*node)

		if (h.direction * currentNode.priority) > (h.direction * parentNode.priority) {
			err = h.array.Update(parentIndex, currentNode)
			if err != nil {
				return err
			}
			err = h.array.Update(currentIndex, parentNode)
			if err != nil {
				return err
			}

			currentIndex = parentIndex
			parentIndex = getParentIndex(currentIndex)
		} else {
			break
		}
	}

	return nil
}

func (h *BinaryHeap) Remove(priority int) error {
	size := h.array.Size()
	if size == 0 {
		return ErrBinaryHeapIsEmpty
	}

	index := 0
	for index < size {
		nodeRaw, err := h.array.Get(index)
		if err != nil {
			return err
		}
		n := nodeRaw.(*node)

		if n.priority == priority {
			// @TODO: need to implement
			return nil
		}
		index++
	}

	return ErrNotFound
}

func (h *BinaryHeap) Update(priority int, newPriority int) error {
	index := 0
	size := h.array.Size()
	for index < size {
		nodeRaw, err := h.array.Get(index)
		if err != nil {
			return err
		}
		n := nodeRaw.(*node)

		if n.priority == priority {
			err = h.array.Update(index, newPriority)
			if err != nil {
				return err
			}
			if (h.direction * newPriority) > (h.direction * priority) {
				err = h.bubbleUp(index)
			} else {
				err = h.bubbleDown(index)
			}
			if err != nil {
				return err
			}
			return nil
		}
		index++
	}

	return ErrNotFound
}

func (h *BinaryHeap) Top() (interface{}, error) {
	if h.array.Size() == 0 {
		return nil, ErrBinaryHeapIsEmpty
	}
	nodeRaw, err := h.array.Get(0)
	if err != nil {
		return nil, err
	}
	n := nodeRaw.(*node)

	if h.array.Size() == 1 {
		err = h.array.Remove(0)
		if err != nil {
			return nil, err
		}
	} else {
		lastNodeIndex := h.array.Size() - 1
		lastNodeRaw, err := h.array.Get(lastNodeIndex)
		if err != nil {
			return nil, err
		}
		lastNode := lastNodeRaw.(*node)
		err = h.array.Update(0, lastNode)
		if err != nil {
			return nil, err
		}
		err = h.array.Remove(lastNodeIndex)
		if err != nil {
			return nil, err
		}
		err = h.bubbleDown(0)
		if err != nil {
			return nil, err
		}
	}

	return n.value, nil
}

func (h *BinaryHeap) Peek() (interface{}, error) {
	if h.array.Size() == 0 {
		return nil, ErrBinaryHeapIsEmpty
	}
	nodeRaw, err := h.array.Get(0)
	if err != nil {
		return nil, err
	}
	node := nodeRaw.(*node)

	return node.value, nil
}

// @TODO: need to add support for multiple same-priority node insertion
func New(elements map[int]interface{}, heapType heapType) (*BinaryHeap, error) {
	direction := 1
	if heapType == Min {
		direction = -1
	}
	s := &BinaryHeap{
		heapType:  heapType,
		direction: direction,
		array:     array.New(nil),
	}

	for priority, value := range elements {
		err := s.Insert(priority, value)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func getParentIndex(index int) int {
	if index%2 != 0 {
		return (index - 1) / 2
	}
	return (index - 2) / 2
}
