package linkedlist

import "fmt"

var ErrInvalidIndex = fmt.Errorf("invalid index")

type node struct {
	value interface{}
	next  *node
}

type LinkedList struct {
	head *node
	size int
}

func (s *LinkedList) Add(value interface{}) {
	n := &node{
		value: value,
	}
	if s.head == nil {
		s.head = n
		s.size = 1
		return
	}
	size := 1
	nextNode := s.head
	for nextNode.next != nil {
		nextNode = nextNode.next
		size++
	}
	nextNode.next = n
	s.size = size + 1
}

func (s *LinkedList) Remove(index int) error {
	if index < 0 || index >= s.Size() {
		return ErrInvalidIndex
	}

	if index == 0 {
		s.head = s.head.next
		s.size--
		return nil
	}

	currentIndex := 0
	currentNode := s.head
	prevNode := s.head
	for currentIndex < index {
		prevNode = currentNode
		currentNode = currentNode.next
		currentIndex++
	}

	prevNode.next = currentNode.next
	s.size--

	return nil
}

func (s *LinkedList) Update(index int, value interface{}) error {
	if index < 0 || index >= s.Size() {
		return ErrInvalidIndex
	}

	currentIndex := 0
	current := s.head
	for currentIndex < index {
		current = current.next
		currentIndex++
	}

	current.value = value
	return nil
}

func (s *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= s.Size() {
		return 0, ErrInvalidIndex
	}

	currentIndex := 0
	current := s.head
	for currentIndex < index {
		current = current.next
		currentIndex++
	}

	return current.value, nil
}

func (s *LinkedList) Size() int {
	return s.size
}

func New(elements []interface{}) *LinkedList {
	s := &LinkedList{}

	for i := range elements {
		s.Add(elements[i])
	}

	return s
}
