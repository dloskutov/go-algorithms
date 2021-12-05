package graph

import (
	"testing"

	"github.com/dloskutov/go-algorithms/src/deprecated/ds"
	"github.com/stretchr/testify/assert"
)

func TestMinimalSpanningTree(t *testing.T) {
	graph, err := ds.NewEdgeGraph([][]int64{
		[]int64{1, 3, 1},
		[]int64{4, 5, 3},
		[]int64{3, 4, 2},
		[]int64{4, 6, 4},
	})
	assert.Nil(t, err)

	edges, err := MinimalSpanningTreePrim(graph)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(edges))

	secondGraph, err := ds.NewEdgeGraph([][]int64{
		[]int64{2, 1, 2},
		[]int64{1, 3, 1},
		[]int64{3, 4, 4},
		[]int64{4, 2, 1},
		[]int64{2, 3, 1},
		[]int64{2, 3, 1},
	})

	edges, err = MinimalSpanningTreePrim(secondGraph)
	assert.Nil(t, err)
	var weight int64
	for _, edge := range edges {
		weight += edge.Weight
	}
	assert.Equal(t, int64(3), weight)
}
