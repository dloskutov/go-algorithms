package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := New([]int{1})

	assert.NotNil(t, a)
}
