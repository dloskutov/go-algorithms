package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New([]int{1})

	assert.NotNil(t, s)
}

func TestEnqueue(t *testing.T) {
	s := New([]int{1})

	s.Enqueue(2)
	s.Enqueue(3)
	s.Enqueue(4)

	assert.Equal(t, 4, s.Size())
}

func TestDequeue(t *testing.T) {
	s := New([]int{1})

	s.Enqueue(2)
	s.Enqueue(3)
	s.Enqueue(4)

	value, err := s.Dequeue()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, value)

	value, err = s.Dequeue()
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, value)

	value, err = s.Dequeue()
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, value)

	value, err = s.Dequeue()
	assert.Equal(t, nil, err)
	assert.Equal(t, 4, value)

	_, err = s.Dequeue()
	assert.Equal(t, ErrEmpty, err)
}
