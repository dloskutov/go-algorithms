package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	h := New(map[string]int{})

	assert.NotNil(t, h)
}
