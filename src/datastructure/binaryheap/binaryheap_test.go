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
