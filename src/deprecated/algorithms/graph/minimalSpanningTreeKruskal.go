package graph

import (
	"github.com/dloskutov/go-algorithms/src/deprecated/ds"
)

// MinimalSpanningTreeKruskal - build minimal spanning tree based on Kruskal's algorithm
func MinimalSpanningTreeKruskal(graph *ds.EdgeGraph) ([]*ds.Edge, error) {
	edges := make([]*ds.Edge, 0)
	edges = append(edges, graph.GetEdges()...)
	vertices := graph.GetVertices()
	MSTEdges := make([]*ds.Edge, 0)
	cc := connectedComponents{
		vertexValueToIndex: make(map[int64]int64),
	}

	for _, vertex := range vertices {
		cc.vertexValueToIndex[vertex.Value] = vertex.Value
	}

	sortEdges(&edges)

	for _, edge := range edges {
		vertexOne := edge.VertexOne
		vertexTwo := edge.VertexTwo
		if !cc.isConnected(vertexOne, vertexTwo) {
			cc.connect(vertexOne, vertexTwo)
			MSTEdges = append(MSTEdges, edge)
		}
	}

	return MSTEdges, nil
}

type connectedComponents struct {
	vertexValueToIndex map[int64]int64
}

func (cc *connectedComponents) isConnected(vertexOne *ds.Vertex, vertexTwo *ds.Vertex) bool {
	vertexOneComponentID := cc.vertexValueToIndex[vertexOne.Value]
	vertexTwoComponentID := cc.vertexValueToIndex[vertexTwo.Value]
	return vertexOneComponentID == vertexTwoComponentID

}

func (cc *connectedComponents) connect(vertexOne *ds.Vertex, vertexTwo *ds.Vertex) {
	vertexOneComponentID := cc.getComponentID(vertexOne)
	vertexTwoComponentID := cc.getComponentID(vertexTwo)
	for vertexValue, vertexValueToIndex := range cc.vertexValueToIndex {
		if vertexValueToIndex == vertexTwoComponentID {
			cc.vertexValueToIndex[vertexValue] = vertexOneComponentID
		}
	}
}

func (cc *connectedComponents) getComponentID(vertex *ds.Vertex) int64 {
	componentID := vertex.Value
	for cc.vertexValueToIndex[componentID] != componentID {
		componentID = cc.vertexValueToIndex[componentID]
	}
	return componentID
}

func sortEdges(edgesToSort *[]*ds.Edge) {
	if len(*edgesToSort) > 0 {
		for {
			needSort := false
			prevEdge := (*edgesToSort)[0]
			for i := 1; i < len(*edgesToSort); i++ {
				currentEdge := (*edgesToSort)[i]
				if currentEdge.Weight < prevEdge.Weight {
					(*edgesToSort)[i-1], (*edgesToSort)[i] = (*edgesToSort)[i], (*edgesToSort)[i-1]
					needSort = true
				}
				prevEdge = currentEdge
			}
			if !needSort {
				break
			}
		}
	}
}
