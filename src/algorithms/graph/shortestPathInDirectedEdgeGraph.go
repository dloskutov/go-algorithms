package graph

import (
	"fmt"
	"math"

	"github.com/dloskutov/go-algorithms/src/ds"
)

// ShortestPathInDirectedEdgeGraph - find the shortest path in directed edge graph
func ShortestPathInDirectedEdgeGraph(graph *ds.DirectedEdgeGraph, fromValue int64, toValue int64) (int64, error) {
	distTo := make(map[*ds.Vertex]int64, 0)
	edgeTo := make(map[*ds.Vertex]*ds.DirectedEdge, 0)
	edgesToExplore := make([]*ds.DirectedEdge, 0)
	vertices := graph.GetVertices()
	var sourceVertex *ds.Vertex
	var distVertex *ds.Vertex
	for _, vertex := range vertices {
		distTo[vertex] = math.MaxInt64
		if vertex.Value == fromValue {
			sourceVertex = vertex
		}
		if vertex.Value == toValue {
			distVertex = vertex
		}
	}
	if sourceVertex == nil {
		return 0, fmt.Errorf("there is no source vertex with value: %d", fromValue)
	}
	if distVertex == nil {
		return 0, fmt.Errorf("there is no dist vertex with value: %d", toValue)
	}
	distTo[sourceVertex] = 0
	firstEdges := graph.GetAllVertexEdges(sourceVertex)
	edgesToExplore = append(edgesToExplore, firstEdges...)

	for len(edgesToExplore) != 0 {
		edgeWithMinWeight := popEdgeWithMinWeight(&edgesToExplore)
		relaxEdge(graph, &edgesToExplore, edgeWithMinWeight, &distTo, &edgeTo)
	}

	return distTo[distVertex], nil
}

func popEdgeWithMinWeight(edgesToExpore *[]*ds.DirectedEdge) *ds.DirectedEdge {
	newEdgesToExplore := make([]*ds.DirectedEdge, 0)

	minEdgeWithWeight := (*edgesToExpore)[0]
	for _, edge := range *edgesToExpore {
		if edge.Weight < minEdgeWithWeight.Weight {
			minEdgeWithWeight = edge
		}
	}

	for _, edge := range *edgesToExpore {
		if edge.Weight != minEdgeWithWeight.Weight {
			newEdgesToExplore = append(newEdgesToExplore, edge)
		}
	}

	(*edgesToExpore) = newEdgesToExplore
	return minEdgeWithWeight
}

func relaxEdge(
	graph *ds.DirectedEdgeGraph,
	edges *[]*ds.DirectedEdge,
	edge *ds.DirectedEdge,
	distTo *map[*ds.Vertex]int64,
	edgeTo *map[*ds.Vertex]*ds.DirectedEdge,
) {
	vertexFrom := edge.VertexFrom
	vertexTo := edge.VertexTo

	if (*distTo)[vertexTo] > edge.Weight+(*distTo)[vertexFrom] {
		(*distTo)[vertexTo] = (*distTo)[vertexFrom] + edge.Weight
		(*edgeTo)[vertexTo] = edge

		newEdges := graph.GetAllVertexEdges(vertexTo)
		(*edges) = append((*edges), newEdges...)
	}
}
