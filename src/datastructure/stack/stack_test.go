package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New([]int{1})

	assert.NotNil(t, s)
}

func TestPut(t *testing.T) {
	s := New([]int{1, 2, 3, 4})

	s.Push(1)
	s.Push(4)

	value, err := s.Pop()
	assert.Equal(t, nil, err)
	assert.Equal(t, 4, value)
}

func TestPop(t *testing.T) {
	s := New(nil)

	value, err := s.Pop()
	assert.Equal(t, ErrEmpty, err)
	assert.Equal(t, 0, value)

	s.Push(1)
	value, err = s.Pop()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, value)

	s.Push(4)
	s.Push(10)
	s.Push(11)

	value, err = s.Pop()
	assert.Equal(t, nil, err)
	assert.Equal(t, 11, value)

	value, err = s.Pop()
	assert.Equal(t, nil, err)
	assert.Equal(t, 10, value)

	value, err = s.Pop()
	assert.Equal(t, nil, err)
	assert.Equal(t, 4, value)

	value, err = s.Pop()
	assert.Equal(t, ErrEmpty, err)
	assert.Equal(t, 0, value)
}
