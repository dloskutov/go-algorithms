package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	bst, err := New(map[int]interface{}{})

	assert.Equal(t, nil, err)
	assert.NotNil(t, bst)
}
