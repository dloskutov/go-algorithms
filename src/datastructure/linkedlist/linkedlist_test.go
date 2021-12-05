package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New([]int{1})

	assert.NotNil(t, s)
}
