package ds

type node struct {
	value int64
	node  *node
}

// LinkedList - linked list data structure
type LinkedList struct {
	first *node
}

// Add - add value to the linked list
func (ll *LinkedList) Add(value int64) {
	ll.first = &node{
		value: value,
		node:  ll.first,
	}
}

// Delete - delete all values from linked list
func (ll *LinkedList) Delete(value int64) {
	for ll.first != nil && ll.first.value == value {
		ll.first = ll.first.node
	}
	if ll.first == nil {
		return
	}

	prevNode := ll.first
	currentNode := prevNode.node
	for currentNode != nil {
		if currentNode.value == value {
			prevNode.node = currentNode.node
		} else {
			prevNode = currentNode
		}
		currentNode = currentNode.node
	}
}

// Contains - check if value exists in linked list
func (ll *LinkedList) Contains(value int64) bool {
	currentNode := ll.first
	for currentNode != nil {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.node
	}
	return false
}

// IsEmpty - check if linked list is empty
func (ll *LinkedList) IsEmpty() bool {
	return ll.first == nil
}
