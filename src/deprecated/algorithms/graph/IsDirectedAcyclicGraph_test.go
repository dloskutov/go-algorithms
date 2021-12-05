package graph

import (
	"testing"

	"github.com/dloskutov/go-algorithms/src/deprecated/ds"
	"github.com/stretchr/testify/assert"
)

func TestIsDirectedAcyclicGraph(t *testing.T) {
	acyclicGraph := ds.NewDirectedGraph([][]int64{
		[]int64{1, 3},
		[]int64{4, 5},
		[]int64{3, 4},
		[]int64{4, 6},
	})
	cyclicGraph := ds.NewDirectedGraph([][]int64{
		[]int64{1, 3},
		[]int64{3, 5},
		[]int64{5, 6},
		[]int64{6, 1},
	})

	assert.True(t, IsDirectedAcyclicGraph(acyclicGraph))
	assert.False(t, IsDirectedAcyclicGraph(cyclicGraph))
}
