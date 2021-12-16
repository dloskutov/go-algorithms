package radixtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	rt, err := New([]string{})

	assert.Equal(t, nil, err)
	assert.NotNil(t, rt)
}
