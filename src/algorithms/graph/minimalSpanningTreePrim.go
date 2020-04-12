package graph

import (
	"fmt"

	"github.com/dloskutov/go-algorithms/src/ds"
)

// MinimalSpanningTreePrim - build minimal spanning tree based on Prim's algorithm
func MinimalSpanningTreePrim(graph *ds.EdgeGraph) ([]*ds.Edge, error) {
	MSTEdges := make([]*ds.Edge, 0)
	edgesToSearch := make([]*ds.Edge, 0)
	vertices := graph.GetVertices()
	addedVertices := make(map[int64]bool)
	checkedEdges := make(map[string]bool)

	if len(vertices) == 0 {
		return nil, fmt.Errorf("edgeGraph is empty")
	}

	vertexEdges := graph.GetAllVertexEdges(vertices[0])
	edgesToSearch = append(edgesToSearch, vertexEdges...)
	addedVertices[vertices[0].Value] = true

	for len(vertices) != len(addedVertices) {
		minWeightedEdge := findMinWieghtedEdge(edgesToSearch)

		if !addedVertices[minWeightedEdge.VertexOne.Value] {
			addEdgesToSearch(&edgesToSearch, checkedEdges, graph.GetAllVertexEdges(minWeightedEdge.VertexOne))
		}
		if !addedVertices[minWeightedEdge.VertexTwo.Value] {
			addEdgesToSearch(&edgesToSearch, checkedEdges, graph.GetAllVertexEdges(minWeightedEdge.VertexTwo))
		}

		if !addedVertices[minWeightedEdge.VertexOne.Value] || !addedVertices[minWeightedEdge.VertexTwo.Value] {
			MSTEdges = append(MSTEdges, minWeightedEdge)
			addedVertices[minWeightedEdge.VertexOne.Value] = true
			addedVertices[minWeightedEdge.VertexTwo.Value] = true
			checkedEdges[fmt.Sprintf("%d_%d", minWeightedEdge.VertexOne.Value, minWeightedEdge.VertexTwo.Value)] = true
			checkedEdges[fmt.Sprintf("%d_%d", minWeightedEdge.VertexTwo.Value, minWeightedEdge.VertexOne.Value)] = true
		}

		removeEdge(&edgesToSearch, minWeightedEdge)
	}

	return MSTEdges, nil
}

func addEdgesToSearch(edgesToSearch *[]*ds.Edge, checkedEdges map[string]bool, edgesToAdd []*ds.Edge) {
	for _, edgeToAdd := range edgesToAdd {
		if !checkedEdges[fmt.Sprintf("%d_%d", edgeToAdd.VertexOne.Value, edgeToAdd.VertexTwo.Value)] {
			hasEdge := false
			for _, edgeToSearch := range *edgesToSearch {
				if edgeToSearch == edgeToAdd {
					hasEdge = true
					break
				}
			}

			if !hasEdge {
				*edgesToSearch = append(*edgesToSearch, edgeToAdd)
			}
		}
	}
}

func removeEdge(edgesToSearch *[]*ds.Edge, edgeToRemove *ds.Edge) {
	newEdgesToSearch := make([]*ds.Edge, 0)

	for _, edge := range *edgesToSearch {
		if edge != edgeToRemove {
			newEdgesToSearch = append(newEdgesToSearch, edge)
		}
	}

	*edgesToSearch = newEdgesToSearch
}

func findMinWieghtedEdge(edges []*ds.Edge) *ds.Edge {
	minWeightedEdge := edges[0]
	for _, edge := range edges {
		if edge.Weight < minWeightedEdge.Weight {
			minWeightedEdge = edge
		}
	}
	return minWeightedEdge
}
