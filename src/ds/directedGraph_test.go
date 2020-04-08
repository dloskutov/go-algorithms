package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectedGraph(t *testing.T) {
	graph := NewDirectedGraph([][]int64{
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

	assert.True(t, graph.HasEdge(1, 2))
	assert.True(t, graph.HasEdge(2, 3))
	assert.True(t, graph.HasEdge(4, 1))
	assert.True(t, graph.HasEdge(1, 4))
	assert.True(t, graph.HasEdge(3, 4))

	assert.False(t, graph.HasEdge(2, 1))
	assert.False(t, graph.HasEdge(4, 3))
	assert.False(t, graph.HasEdge(4, 2))
	assert.False(t, graph.HasEdge(4, 4))
}
