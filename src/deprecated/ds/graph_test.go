package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph(t *testing.T) {
	graph := NewGraph([][]int64{
		[]int64{1, 2},
		[]int64{5, 4},
		[]int64{2, 3},
		[]int64{3, 4},
		[]int64{4, 1},
		[]int64{2, 4},
		[]int64{1, 4},
	})

	assert.True(t, graph.HasVertexWithValue(1))
	assert.True(t, graph.HasVertexWithValue(2))
	assert.True(t, graph.HasVertexWithValue(3))
	assert.True(t, graph.HasVertexWithValue(4))
	assert.True(t, graph.HasVertexWithValue(5))

	assert.False(t, graph.HasVertexWithValue(8))
	assert.False(t, graph.HasVertexWithValue(6))
}
