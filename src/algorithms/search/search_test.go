package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSequentialSearch(t *testing.T) {
	sequentialSearch := new(SequentialSearch)

	value, err := sequentialSearch.Get("non-exist key")
	assert.Equal(t, int64(0), value)
	assert.Error(t, err)

	sequentialSearch.Put("first", 1)
	sequentialSearch.Put("second", 2)
	sequentialSearch.Put("third", 3)

	value, err = sequentialSearch.Get("second")
	assert.Equal(t, int64(2), value)
	assert.Nil(t, err)

	sequentialSearch.Put("second", 20)
	value, err = sequentialSearch.Get("second")
	assert.Equal(t, int64(20), value)
	assert.Nil(t, err)
}
