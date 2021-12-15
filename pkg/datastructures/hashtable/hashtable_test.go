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

	err = h.Put("other", "value_3")
	assert.Equal(t, nil, err)

	err = h.Put("key", "value_1")
	assert.Equal(t, nil, err)

	err = h.Put("yek", "value_2")
	assert.Equal(t, nil, err)

	value, err := h.Get("yek")
	assert.Equal(t, nil, err)
	assert.Equal(t, "value_2", value)

	_, err = h.Get("non_exists")
	assert.Equal(t, ErrInvalidKey, err)
}

func TestUpdate(t *testing.T) {
	h, err := New(map[string]interface{}{})

	assert.Equal(t, nil, err)

	err = h.Put("key", "value_1")
	assert.Equal(t, nil, err)

	err = h.Update("key", "value_2")
	assert.Equal(t, nil, err)

	value, err := h.Get("key")
	assert.Equal(t, nil, err)
	assert.Equal(t, "value_2", value)

	err = h.Update("non_exists", "empty")
	assert.Equal(t, ErrInvalidKey, err)
}

func TestRemove(t *testing.T) {
	h, err := New(map[string]interface{}{})

	assert.Equal(t, nil, err)

	err = h.Put("key", "value_1")
	assert.Equal(t, nil, err)

	err = h.Remove("key")
	assert.Equal(t, nil, err)

	err = h.Remove("key")
	assert.Equal(t, ErrInvalidKey, err)
}
