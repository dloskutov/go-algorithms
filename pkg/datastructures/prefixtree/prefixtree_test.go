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

func TestContains(t *testing.T) {
	pt, err := New(map[string]interface{}{
		"first":  1,
		"second": 2,
		"third":  3,
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, true, pt.Contains("first"))
	assert.Equal(t, false, pt.Contains("firs"))
	assert.Equal(t, false, pt.Contains("other"))
}

func TestRemove(t *testing.T) {
	pt, err := New(map[string]interface{}{
		"first":  1,
		"second": 2,
		"third":  3,
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, true, pt.Contains("first"))

	assert.Equal(t, ErrInvalidKey, pt.Remove(""))
	assert.Equal(t, ErrInvalidKey, pt.Remove("other"))
	assert.Equal(t, nil, pt.Remove("first"))

	assert.Equal(t, false, pt.Contains("first"))
}
