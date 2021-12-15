package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New([]interface{}{1})

	assert.NotNil(t, s)
}

func TestSize(t *testing.T) {
	s := New([]interface{}{1, 2, 3})

	assert.Equal(t, 3, s.Size())
	s.Add(5)
	assert.Equal(t, 4, s.Size())
}

func TestGet(t *testing.T) {
	s := New([]interface{}{1, 2, 3})

	value, err := s.Get(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, value)

	value, err = s.Get(2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, value)

	_, err = s.Get(-1)
	assert.Equal(t, ErrInvalidIndex, err)

	_, err = s.Get(100)
	assert.Equal(t, ErrInvalidIndex, err)
}

func TestRemove(t *testing.T) {
	s := New([]interface{}{1, 2, 3})

	err := s.Remove(1)
	assert.Equal(t, nil, err)

	value, err := s.Get(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, value)

	err = s.Remove(0)
	assert.Equal(t, nil, err)

	err = s.Remove(-1)
	assert.Equal(t, ErrInvalidIndex, err)

	err = s.Remove(100)
	assert.Equal(t, ErrInvalidIndex, err)
}

func TestUpdate(t *testing.T) {
	s := New([]interface{}{1, 2, 3})

	err := s.Update(1, 11)
	assert.Equal(t, nil, err)

	value, err := s.Get(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 11, value)

	err = s.Update(-1, 12)
	assert.Equal(t, ErrInvalidIndex, err)

	err = s.Update(100, 10)
	assert.Equal(t, ErrInvalidIndex, err)
}
