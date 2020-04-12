package ds

import "fmt"

// EdgeGraph - graph based on edges data strcture
type EdgeGraph struct {
	edges []*Edge
}

// Edge - edge of EdgeGraph
type Edge struct {
	VertexOne *Vertex
	VertexTwo *Vertex
	Weight    int64
}

// NewEdgeGraph - create edge graph based on edges
func NewEdgeGraph(edges [][]int64) (*EdgeGraph, error) {
	graph := &EdgeGraph{
		edges: make([]*Edge, 0),
	}

	for _, edge := range edges {
		if len(edge) != 3 {
			return nil, fmt.Errorf("each array for edge must have 3 paramas")
		}
		vertexOne := graph.GetVertexByValue(edge[0])
		vertexTwo := graph.GetVertexByValue(edge[1])

		if vertexOne == nil {
			vertexOne = &Vertex{
				Value: edge[0],
			}
		}
		if vertexTwo == nil {
			vertexTwo = &Vertex{
				Value: edge[1],
			}
		}

		graph.addEdge(vertexOne, vertexTwo, edge[2])
	}
	return graph, nil
}

// GetAllVertexEdges - get all edges connected to vertex
func (g *EdgeGraph) GetAllVertexEdges(vertex *Vertex) []*Edge {
	edges := make([]*Edge, 0)
	for _, edge := range g.edges {
		if edge.VertexOne.Value == vertex.Value || edge.VertexTwo.Value == vertex.Value {
			edges = append(edges, edge)
		}
	}
	return edges
}

// GetVertices - get all vertices
func (g *EdgeGraph) GetVertices() []*Vertex {
	vertices := make([]*Vertex, 0)
	values := make(map[int64]bool)
	for _, edge := range g.edges {
		if !values[edge.VertexOne.Value] {
			vertices = append(vertices, edge.VertexOne)
			values[edge.VertexOne.Value] = true
		}
		if !values[edge.VertexTwo.Value] {
			vertices = append(vertices, edge.VertexTwo)
			values[edge.VertexTwo.Value] = true
		}
	}
	return vertices
}

// GetEdges - get all edges
func (g *EdgeGraph) GetEdges() []*Edge {
	return g.edges
}

// GetVertexByValue - get vertex by value
func (g *EdgeGraph) GetVertexByValue(value int64) *Vertex {
	for _, edge := range g.edges {
		if edge.VertexOne.Value == value {
			return edge.VertexOne
		}
		if edge.VertexTwo.Value == value {
			return edge.VertexTwo
		}
	}
	return nil
}

// GetEdge - get edge by values of edges
func (g *EdgeGraph) GetEdge(vertexOneValue int64, vertexTwoValue int64) (*Edge, error) {
	for _, edge := range g.edges {
		if edge.VertexOne.Value == vertexOneValue && edge.VertexTwo.Value == vertexTwoValue {
			return edge, nil
		}
	}
	return nil, fmt.Errorf("unable to find edge with '%d' and '%d'  vertices", vertexOneValue, vertexTwoValue)
}

func (g *EdgeGraph) addEdge(vertexOne *Vertex, vertexTwo *Vertex, weight int64) {
	for _, edge := range g.edges {
		if edge.VertexOne == vertexOne && edge.VertexTwo == vertexTwo {
			return
		}
		if edge.VertexTwo == vertexOne && edge.VertexOne == vertexOne {
			return
		}
	}
	g.edges = append(g.edges, &Edge{
		VertexOne: vertexOne,
		VertexTwo: vertexTwo,
		Weight:    weight,
	})
}
