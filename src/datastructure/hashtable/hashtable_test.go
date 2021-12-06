package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	h, err := New(map[string]interface{}{
		"key": "value",
	})

	assert.NotNil(t, h)
	assert.Equal(t, nil, err)
}

func TestPut(t *testing.T) {
	h, err := New(map[string]interface{}{})

	assert.Equal(t, nil, err)

	err = h.Put("key", "value")
	assert.Equal(t, nil, err)

	value, err := h.Get("key")
	assert.Equal(t, nil, err)
	assert.Equal(t, "value", value)
}
