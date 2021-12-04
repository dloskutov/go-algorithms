package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	h := New(nil, Max)

	assert.NotNil(t, h)
}
