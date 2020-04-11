package graph

import (
	"fmt"

	"github.com/dloskutov/go-algorithms/src/ds"
)

// TopologicalSort - return verices of directed graph in topological sort
func TopologicalSort(graph *ds.DirectedGraph) ([]*ds.Vertex, error) {
	visitedVertices := make(map[int64]bool)
	stack := make([]*ds.Vertex, 0)

	if !IsDirectedAcyclicGraph(graph) {
		return nil, fmt.Errorf("there is cycle in graph")
	}

	for _, v := range graph.GetVertices() {
		visitedVertices[v.Value] = false
	}

	for _, vertexToVisit := range graph.GetVertices() {
		if !visitedVertices[vertexToVisit.Value] {
			topologicalDFS(vertexToVisit, &visitedVertices, &stack)
		}
	}

	return stack, nil
}

func topologicalDFS(
	nextVertex *ds.Vertex,
	visitedVertices *map[int64]bool,
	stack *[]*ds.Vertex,
) {
	(*visitedVertices)[nextVertex.Value] = true

	for _, nextVertex := range nextVertex.AdjacentVertices {
		if !(*visitedVertices)[nextVertex.Value] {
			topologicalDFS(nextVertex, visitedVertices, stack)
		}
	}

	(*stack) = append(*stack, nextVertex)
}
