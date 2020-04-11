package graph

import (
	"testing"

	"github.com/dloskutov/go-algorithms/src/ds"
	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort(t *testing.T) {
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

	expectedVerticesValues := []int64{5, 6, 4, 3, 1}

	sortedVertices, err := TopologicalSort(acyclicGraph)
	assert.Nil(t, err)

	for index, expectedVerticesValue := range expectedVerticesValues {
		sortedValue := sortedVertices[index].Value
		assert.Equal(t, sortedValue, expectedVerticesValue)
	}

	_, err = TopologicalSort(cyclicGraph)
	assert.NotNil(t, err)
}
