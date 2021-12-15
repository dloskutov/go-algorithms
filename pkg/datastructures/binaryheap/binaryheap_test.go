package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	h, err := New(nil, Max)

	assert.Nil(t, err)
	assert.NotNil(t, h)
}

func TestInsert_Max(t *testing.T) {
	h, err := New(map[int]interface{}{
		1:  "first",
		21: "seventh",
		11: "fourth",
		22: "fifth",
		2:  "second",
		10: "third",
		13: "sixth",
	}, Max)
	assert.Equal(t, nil, err)

	value, err := h.Peek()
	assert.Equal(t, nil, err)
	assert.Equal(t, "fifth", value)
}

func TestInsert_Min(t *testing.T) {
	h, err := New(map[int]interface{}{
		1:  "first",
		21: "seventh",
		11: "fourth",
		22: "fifth",
		2:  "second",
		10: "third",
		13: "sixth",
	}, Min)
	assert.Equal(t, nil, err)

	value, err := h.Peek()
	assert.Equal(t, nil, err)
	assert.Equal(t, "first", value)
}

func TestPeek(t *testing.T) {
	h, err := New(nil, Max)
	assert.Equal(t, nil, err)

	_, err = h.Peek()
	assert.Equal(t, ErrBinaryHeapIsEmpty, err)
}

func TestTop(t *testing.T) {
	h, err := New(nil, Max)
	assert.Equal(t, nil, err)

	_, err = h.Top()
	assert.Equal(t, ErrBinaryHeapIsEmpty, err)

	h, err = New(map[int]interface{}{
		1:  "first",
		21: "seventh",
		11: "fourth",
		22: "fifth",
		2:  "second",
		10: "third",
		13: "sixth",
	}, Max)
	assert.Equal(t, nil, err)

	items := []string{"fifth", "seventh", "sixth", "fourth", "third", "second", "first"}
	for i := range items {
		value, err := h.Top()
		assert.Equal(t, nil, err)
		assert.Equal(t, value, items[i])
	}
}

func TestUpdate(t *testing.T) {
	h, err := New(nil, Max)
	assert.Equal(t, nil, err)

	_, err = h.Top()
	assert.Equal(t, ErrBinaryHeapIsEmpty, err)

	h, err = New(map[int]interface{}{
		1:  "first",
		21: "seventh",
		11: "fourth",
		22: "fifth",
		2:  "second",
		10: "third",
		13: "sixth",
	}, Max)
	assert.Equal(t, nil, err)

	err = h.Update(33, 44)
	assert.Equal(t, ErrNotFound, err)

	items := []string{"fifth", "seventh", "sixth", "fourth", "third", "second", "first"}
	for i := range items {
		value, err := h.Top()
		assert.Equal(t, nil, err)
		assert.Equal(t, value, items[i])
	}

	h, err = New(map[int]interface{}{
		1:  "first",
		21: "seventh",
		11: "fourth",
		22: "fifth",
		2:  "second",
		10: "third",
		13: "sixth",
	}, Max)
	assert.Equal(t, nil, err)

	err = h.Update(11, 33)
	assert.Equal(t, nil, err)

	items = []string{"fourth", "fifth", "seventh", "sixth", "third", "second", "first"}
	for i := range items {
		value, err := h.Top()
		assert.Equal(t, nil, err)
		assert.Equal(t, value, items[i])
	}
}

func TestRemove(t *testing.T) {
	h, err := New(nil, Max)

	assert.Equal(t, nil, err)

	err = h.Remove(0)
	assert.Equal(t, ErrBinaryHeapIsEmpty, err)

	err = h.Insert(1, 10)
	assert.Equal(t, nil, err)
	err = h.Insert(2, 20)
	assert.Equal(t, nil, err)

	err = h.Remove(10)
	assert.Equal(t, ErrNotFound, err)

	h, err = New(map[int]interface{}{
		1:  "first",
		21: "seventh",
		11: "fourth",
		22: "fifth",
		2:  "second",
		10: "third",
		13: "sixth",
	}, Max)
	assert.Equal(t, nil, err)

	err = h.Remove(11)
	assert.Equal(t, nil, err)

	err = h.Remove(1)
	assert.Equal(t, nil, err)

	err = h.Remove(22)
	assert.Equal(t, nil, err)

	items := []string{"seventh", "sixth", "third", "second"}
	for i := range items {
		value, err := h.Top()
		assert.Equal(t, nil, err)
		assert.Equal(t, value, items[i])
	}

	h, err = New(map[int]interface{}{
		1: "first",
		2: "second",
	}, Max)
	assert.Equal(t, nil, err)

	err = h.Remove(2)
	assert.Equal(t, nil, err)

	err = h.Remove(1)
	assert.Equal(t, nil, err)
}
