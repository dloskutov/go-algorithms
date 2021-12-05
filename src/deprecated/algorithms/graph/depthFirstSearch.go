package graph

import (
	"github.com/dloskutov/go-algorithms/src/deprecated/ds"
)

// DepthFirstSearch - check if there is path
// from vertex with sourceValue to vertex with checkValue
func DepthFirstSearch(graph *ds.Graph, sourceValue int64, checkValue int64) bool {
	marked := make(map[int64]bool, 0)

	sourceVertex := graph.GetVertexByValue(sourceValue)
	if sourceVertex != nil {
		depthFirstSearch(marked, sourceVertex)

		return marked[checkValue]
	}
	return false
}

func depthFirstSearch(marked map[int64]bool, sourceVertex *ds.Vertex) {
	for _, vertex := range sourceVertex.AdjacentVertices {
		if !marked[vertex.Value] {
			marked[vertex.Value] = true
			depthFirstSearch(marked, vertex)
		}
	}
}
