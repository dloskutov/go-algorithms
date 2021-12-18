package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	bst, err := New(map[int]interface{}{
		10: "first",
		2:  "second",
		5:  "third",
		7:  "forth",
		12: "fifth",
	})

	assert.Equal(t, nil, err)
	assert.NotNil(t, bst)
}

func TestSearch(t *testing.T) {
	bst, err := New(map[int]interface{}{
		10: "first",
		2:  "second",
		5:  "third",
		7:  "forth",
		12: "fifth",
	})

	assert.Equal(t, nil, err)
	assert.NotNil(t, bst)

	value, err := bst.Search(5)
	assert.Equal(t, nil, err)
	assert.Equal(t, value, "third")

	_, err = bst.Search(111)
	assert.Equal(t, ErrInvalidKey, err)
}
