package bloomfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	b := New()

	assert.NotNil(t, b)
}

func TestInsert(t *testing.T) {
	b := New()

	b.Insert("key")
	assert.Equal(t, true, b.Contains("key"))
	assert.Equal(t, false, b.Contains("yek"))
}

func TestContains(t *testing.T) {
	b := New()

	assert.Equal(t, false, b.Contains("key"))
	b.Insert("key")
	assert.Equal(t, true, b.Contains("key"))
}
