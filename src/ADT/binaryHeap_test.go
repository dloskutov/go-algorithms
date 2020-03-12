package adt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryHeap(t *testing.T) {
	array := []int64{4, 1, 6, 3, 8, 4, 1, 9, 7, 11, 4, 33, 23}
	heap := CreateBinaryHeap(array)

	assert.Equal(t, false, heap.IsEmpty())

	prevValue := int64(math.MaxInt64)
	for range array {
		value := heap.Get()
		t.Log(value)
		if value <= prevValue {
			prevValue = value
		} else {
			assert.Fail(t, "value")
		}
	}

	assert.Equal(t, true, heap.IsEmpty())
	assert.Equal(t, 0, heap.Get())
}
