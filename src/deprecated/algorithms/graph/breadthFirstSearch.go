package graph

import (
	"github.com/dloskutov/go-algorithms/src/deprecated/ds"
)

// BreadthFirstSearch - check if there is path
// from vertex with sourceValue to vertex with checkValue
func BreadthFirstSearch(graph *ds.Graph, sourceValue int64, checkValue int64) bool {
	marked := make(map[int64]bool, 0)

	sourceVertex := graph.GetVertexByValue(sourceValue)
	if sourceVertex != nil {
		verticesToCheck := make([]*ds.Vertex, 0)
		verticesToCheck = append(verticesToCheck, sourceVertex.AdjacentVertices...)

		i := 0
		for i < len(verticesToCheck) {
			vertexToCheck := verticesToCheck[i]
			if !marked[vertexToCheck.Value] {
				marked[vertexToCheck.Value] = true
				verticesToCheck = append(verticesToCheck, vertexToCheck.AdjacentVertices...)
			}
			i++
		}

		return marked[checkValue]
	}
	return false
}
