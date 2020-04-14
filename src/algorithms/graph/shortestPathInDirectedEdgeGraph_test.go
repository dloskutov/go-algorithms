package graph

import (
	"testing"

	"github.com/dloskutov/go-algorithms/src/ds"
	"github.com/stretchr/testify/assert"
)

func TestShortestPathInDirectedEdgeGraph(t *testing.T) {
	graph, err := ds.NewDirectedEdgeGraph([][]int64{
		[]int64{1, 3, 50},
		[]int64{4, 5, 40},
		[]int64{3, 4, 30},
		[]int64{4, 6, 40},
		[]int64{1, 6, 10},
		[]int64{6, 4, 2},
		[]int64{4, 6, 4},
	})
	assert.Nil(t, err)

	pathLength, err := ShortestPathInDirectedEdgeGraph(graph, 1, 4)
	assert.Nil(t, err)
	assert.Equal(t, int64(12), pathLength)
}
