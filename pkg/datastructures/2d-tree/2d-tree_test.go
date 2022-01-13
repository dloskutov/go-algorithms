package twodtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New([][2]int{{}})

	assert.NotNil(t, s)
}
