package prefixtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	pt, err := New(map[string]interface{}{})

	assert.Equal(t, nil, err)
	assert.NotNil(t, pt)
}
