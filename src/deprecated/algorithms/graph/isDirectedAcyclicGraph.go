package graph

import (
	"github.com/dloskutov/go-algorithms/src/deprecated/ds"
)

// IsDirectedAcyclicGraph - check if direct graph has no cycles
func IsDirectedAcyclicGraph(graph *ds.DirectedGraph) bool {
	notVisitedVertices := make(map[int64]bool, 0)
	inProgressVertices := make(map[int64]bool, 0)
	visitedVertices := make(map[int64]bool, 0)

	for _, vertex := range graph.GetVertices() {
		notVisitedVertices[vertex.Value] = true
	}

	hasCycle := false
	for {
		nextVertex := getNextVertex(graph.GetVertices(), notVisitedVertices)
		if nextVertex == nil {
			break
		}
		if dfs(nextVertex, &notVisitedVertices, &inProgressVertices, &visitedVertices) {
			hasCycle = true
			break
		}
	}

	return !hasCycle
}

func getNextVertex(vertices []*ds.Vertex, notVisitedVertices map[int64]bool) *ds.Vertex {
	for _, vertex := range vertices {
		if notVisitedVertices[vertex.Value] {
			return vertex
		}
	}
	return nil
}

func dfs(
	nextVertex *ds.Vertex,
	notVisitedVertices *map[int64]bool,
	inProgressVertices *map[int64]bool,
	visitedVertices *map[int64]bool,
) bool {
	if (*visitedVertices)[nextVertex.Value] {
		return false
	}
	if (*inProgressVertices)[nextVertex.Value] {
		return true
	}
	(*notVisitedVertices)[nextVertex.Value] = false
	(*inProgressVertices)[nextVertex.Value] = true

	for _, vertex := range nextVertex.AdjacentVertices {
		if dfs(vertex, notVisitedVertices, inProgressVertices, visitedVertices) {
			return true
		}
	}

	(*inProgressVertices)[nextVertex.Value] = false
	(*visitedVertices)[nextVertex.Value] = true

	return false
}
