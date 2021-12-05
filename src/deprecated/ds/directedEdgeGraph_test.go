package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectedEdgeGraph(t *testing.T) {
	graph, err := NewDirectedEdgeGraph([][]int64{
		[]int64{1, 2, 4},
		[]int64{5, 4, 3},
		[]int64{2, 3, 5},
		[]int64{3, 4, 2},
		[]int64{4, 1, 6},
		[]int64{4, 1, 6},
		[]int64{2, 4, 2},
		[]int64{1, 4, 1},
	})
	assert.Nil(t, err)

	assert.Equal(t, 7, len(graph.GetEdges()))

	vertices := graph.GetVertices()
	edges := graph.GetAllVertexEdges(vertices[0])
	assert.Equal(t, 2, len(edges))
}
