package binaryheap

import (
	"fmt"

	"github.com/dloskutov/go-algorithms/src/datastructure/array"
)

var ErrBinaryHeapIsEmpty = fmt.Errorf("binary heap is empty")

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
	// @TODO: need to implement
	return nil
}

func (h *BinaryHeap) Update(priority int, newPriority int) error {
	// @TODO: need to implement
	return nil
}

func (h *BinaryHeap) Top() (interface{}, error) {
	// @TODO: need to implement
	return nil, nil
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
