package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEdgeGraph(t *testing.T) {
	graph, err := NewEdgeGraph([][]int64{
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

	edge, err := graph.GetEdge(3, 4)
	assert.Nil(t, err)
	assert.Equal(t, int64(3), edge.VertexFrom.Value)
	assert.Equal(t, int64(4), edge.VertexTo.Value)

	assert.Equal(t, 7, len(graph.GetEdges()))

	_, err = graph.GetEdge(4, 11)
	assert.NotNil(t, err)
}
