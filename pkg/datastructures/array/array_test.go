package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New([]interface{}{1})

	assert.NotNil(t, s)
}

func TestGet(t *testing.T) {
	s := New([]interface{}{1, 2, 3, 4})

	value, err := s.Get(2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, value)

	_, err = s.Get(-1)
	assert.Equal(t, ErrInvalidIndex, err)

	_, err = s.Get(100)
	assert.Equal(t, ErrInvalidIndex, err)
}

func TestSize(t *testing.T) {
	s := New([]interface{}{1, 2, 3, 4})

	assert.Equal(t, 4, s.Size())
}

func TestRemove(t *testing.T) {
	s := New([]interface{}{1, 2, 3, 4})

	err := s.Remove(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, s.Size())

	value, err := s.Get(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, value)

	err = s.Remove(-1)
	assert.Equal(t, ErrInvalidIndex, err)

	err = s.Remove(100)
	assert.Equal(t, ErrInvalidIndex, err)
}

func TestUpdate(t *testing.T) {
	s := New([]interface{}{1, 2, 3, 4})

	err := s.Update(1, 10)
	assert.Equal(t, nil, err)

	value, err := s.Get(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 10, value)

	err = s.Update(-1, 11)
	assert.Equal(t, ErrInvalidIndex, err)

	err = s.Update(100, 12)
	assert.Equal(t, ErrInvalidIndex, err)
}

func TestAdd(t *testing.T) {
	s := New([]interface{}{1, 2, 3})

	s.Add(4)

	value, err := s.Get(3)
	assert.Equal(t, nil, err)
	assert.Equal(t, 4, value)
}
