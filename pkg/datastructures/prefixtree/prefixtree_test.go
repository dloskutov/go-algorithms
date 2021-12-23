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

func TestGet(t *testing.T) {
	pt, err := New(map[string]interface{}{
		"first":  1,
		"second": 2,
		"third":  3,
	})

	assert.Equal(t, nil, err)

	value, err := pt.Get("first")
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, value)

	_, err = pt.Get("fir")
	assert.Equal(t, ErrInvalidKey, err)

	_, err = pt.Get("")
	assert.Equal(t, ErrInvalidKey, err)

	_, err = pt.Get("invalid key")
	assert.Equal(t, ErrInvalidKey, err)
}
