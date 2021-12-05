package ds

import "fmt"

// DirectedEdgeGraph - graph based on edges data structure
type DirectedEdgeGraph struct {
	edges []*DirectedEdge
}

// DirectedEdge - edge of DirectedEdgeGraph
type DirectedEdge struct {
	VertexFrom *Vertex
	VertexTo   *Vertex
	Weight     int64
}

// NewDirectedEdgeGraph - create edge graph based on edges
func NewDirectedEdgeGraph(edges [][]int64) (*DirectedEdgeGraph, error) {
	graph := &DirectedEdgeGraph{
		edges: make([]*DirectedEdge, 0),
	}

	for _, edge := range edges {
		if len(edge) != 3 {
			return nil, fmt.Errorf("each array for edge must have 3 paramas")
		}
		vertexFrom := graph.GetVertexByValue(edge[0])
		vertexTo := graph.GetVertexByValue(edge[1])

		if vertexFrom == nil {
			vertexFrom = &Vertex{
				Value: edge[0],
			}
		}
		if vertexTo == nil {
			vertexTo = &Vertex{
				Value: edge[1],
			}
		}

		graph.addEdge(vertexFrom, vertexTo, edge[2])
	}
	return graph, nil
}

// GetAllVertexEdges - get all directed edges from vertex
func (g *DirectedEdgeGraph) GetAllVertexEdges(vertex *Vertex) []*DirectedEdge {
	edges := make([]*DirectedEdge, 0)
	for _, edge := range g.edges {
		if edge.VertexFrom.Value == vertex.Value {
			edges = append(edges, edge)
		}
	}
	return edges
}

// GetVertices - get all vertices
func (g *DirectedEdgeGraph) GetVertices() []*Vertex {
	vertices := make([]*Vertex, 0)
	values := make(map[int64]bool)
	for _, edge := range g.edges {
		if !values[edge.VertexFrom.Value] {
			vertices = append(vertices, edge.VertexFrom)
			values[edge.VertexFrom.Value] = true
		}
		if !values[edge.VertexTo.Value] {
			vertices = append(vertices, edge.VertexTo)
			values[edge.VertexTo.Value] = true
		}
	}
	return vertices
}

// GetEdges - get all edges
func (g *DirectedEdgeGraph) GetEdges() []*DirectedEdge {
	return g.edges
}

// GetVertexByValue - get vertex by value
func (g *DirectedEdgeGraph) GetVertexByValue(value int64) *Vertex {
	for _, edge := range g.edges {
		if edge.VertexFrom.Value == value {
			return edge.VertexFrom
		}
		if edge.VertexTo.Value == value {
			return edge.VertexTo
		}
	}
	return nil
}

func (g *DirectedEdgeGraph) addEdge(vertexFrom *Vertex, vertexTo *Vertex, weight int64) {
	for _, edge := range g.edges {
		if edge.VertexFrom == vertexFrom && edge.VertexTo == vertexTo {
			return
		}
	}
	g.edges = append(g.edges, &DirectedEdge{
		VertexFrom: vertexFrom,
		VertexTo:   vertexTo,
		Weight:     weight,
	})
}
