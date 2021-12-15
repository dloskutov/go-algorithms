package disjointset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	d := New([]int{})

	assert.NotNil(t, d)
}

func TestAdd(t *testing.T) {
	d := New([]int{1, 2, 3, 4, 5})

	assert.Equal(t, true, d.Add(6))
	assert.Equal(t, true, d.Add(7))
	assert.Equal(t, false, d.Add(1))
	assert.Equal(t, false, d.Add(3))
}

func TestFindPartition(t *testing.T) {
	d := New([]int{1, 2, 3, 4, 5})

	assert.Equal(t, 1, d.FindPartition(1))
	assert.Equal(t, 2, d.FindPartition(2))
	assert.Equal(t, 3, d.FindPartition(3))
	assert.Equal(t, 4, d.FindPartition(4))
	assert.Equal(t, 5, d.FindPartition(5))

	d.Merge(1, 3)
	d.Merge(2, 4)
	d.Merge(3, 5)

	assert.Equal(t, d.FindPartition(3), d.FindPartition(1))
	assert.Equal(t, d.FindPartition(2), d.FindPartition(4))
	assert.Equal(t, d.FindPartition(3), d.FindPartition(5))
}

func TestAreDisjoint(t *testing.T) {
	d := New([]int{1, 2, 3, 4, 5})

	d.Merge(1, 3)
	d.Merge(2, 4)
	d.Merge(3, 5)

	assert.Equal(t, false, d.AreDisjoint(1, 3))
	assert.Equal(t, false, d.AreDisjoint(3, 5))
	assert.Equal(t, false, d.AreDisjoint(2, 4))
	assert.Equal(t, true, d.AreDisjoint(3, 4))
	assert.Equal(t, true, d.AreDisjoint(2, 5))
	assert.Equal(t, true, d.AreDisjoint(1, 4))
}
