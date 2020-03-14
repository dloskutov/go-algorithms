package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	linkedList := new(LinkedList)

	assert.True(t, linkedList.IsEmpty())

	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)

	assert.False(t, linkedList.IsEmpty())

	assert.True(t, linkedList.Contains(1))
	assert.True(t, linkedList.Contains(2))
	assert.True(t, linkedList.Contains(3))

	assert.False(t, linkedList.Contains(10))
	assert.False(t, linkedList.Contains(20))
	assert.False(t, linkedList.Contains(30))

	linkedList.Delete(3)
	linkedList.Delete(1)
	linkedList.Delete(2)

	assert.True(t, linkedList.IsEmpty())

	assert.False(t, linkedList.Contains(1))
	assert.False(t, linkedList.Contains(2))
	assert.False(t, linkedList.Contains(3))

}
