package graph

import (
	"testing"

	"github.com/dloskutov/go-algorithms/src/deprecated/ds"
	"github.com/stretchr/testify/assert"
)

func TestDepthFirstSearch(t *testing.T) {
	graph := ds.NewGraph([][]int64{
		[]int64{1, 3},
		[]int64{4, 5},
		[]int64{3, 4},
		[]int64{4, 5},
	})

	assert.True(t, DepthFirstSearch(graph, 1, 3))
	assert.True(t, DepthFirstSearch(graph, 3, 5))
	assert.True(t, DepthFirstSearch(graph, 1, 5))

	assert.False(t, DepthFirstSearch(graph, 1, 6))
	assert.False(t, DepthFirstSearch(graph, 3, 7))
	assert.False(t, DepthFirstSearch(graph, 8, 1))
}
