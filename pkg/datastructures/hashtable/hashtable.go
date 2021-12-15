package hashtable

import (
	"fmt"

	"github.com/dloskutov/go-algorithms/pkg/datastructures/array"
	"github.com/dloskutov/go-algorithms/pkg/datastructures/linkedlist"
)

var ErrInvalidKey = fmt.Errorf("invalid key")

const hashMod = 127

type node struct {
	key   string
	value interface{}
}

type structure struct {
	table *array.Array
}

func (s *structure) Put(key string, value interface{}) error {
	keyIndex := hash(key)
	linkedList, err := s.table.Get(keyIndex)
	if err != nil {
		return ErrInvalidKey
	}
	linkedList.(*linkedlist.LinkedList).Add(&node{
		key:   key,
		value: value,
	})

	return nil
}

func (s *structure) Remove(key string) error {
	keyIndex := hash(key)
	linkedListRaw, err := s.table.Get(keyIndex)
	if err != nil {
		return ErrInvalidKey
	}
	linkedList := linkedListRaw.(*linkedlist.LinkedList)

	size := linkedList.Size()
	index := 0
	for index < size {
		nodeRaw, err := linkedList.Get(index)
		if err != nil {
			return ErrInvalidKey
		}
		n := nodeRaw.(*node)
		if n.key == key {
			return linkedList.Remove(index)
		}
		index++
	}

	return ErrInvalidKey
}

func (s *structure) Update(key string, value interface{}) error {
	keyIndex := hash(key)
	linkedListRaw, err := s.table.Get(keyIndex)
	if err != nil {
		return ErrInvalidKey
	}
	linkedList := linkedListRaw.(*linkedlist.LinkedList)

	size := linkedList.Size()
	index := 0
	for index < size {
		nodeRaw, err := linkedList.Get(index)
		if err != nil {
			return ErrInvalidKey
		}
		n := nodeRaw.(*node)
		if n.key == key {
			return linkedList.Update(index, &node{
				key:   n.key,
				value: value,
			})
		}
		index++
	}

	return ErrInvalidKey
}

func (s *structure) Get(key string) (interface{}, error) {
	keyIndex := hash(key)
	linkedListRaw, err := s.table.Get(keyIndex)
	if err != nil {
		return 0, ErrInvalidKey
	}
	linkedList := linkedListRaw.(*linkedlist.LinkedList)

	size := linkedList.Size()
	index := 0
	for index < size {
		nodeRaw, err := linkedList.Get(index)
		if err != nil {
			return 0, ErrInvalidKey
		}
		n := nodeRaw.(*node)
		if n.key == key {
			return n.value, nil
		}
		index++
	}

	return 0, ErrInvalidKey
}

func New(keyValues map[string]interface{}) (*structure, error) {
	s := &structure{
		table: array.New(nil),
	}

	index := 0
	for index < hashMod {
		index++
		s.table.Add(linkedlist.New(nil))
	}

	for key, value := range keyValues {
		err := s.Put(key, value)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func hash(key string) int {
	value := 0

	for index, r := range key {
		value += int(r) + int(int(r)/(index+1))
	}

	return value % hashMod
}
