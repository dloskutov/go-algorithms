package ds

import "fmt"

// EdgeGraph - graph based on edges data strcture
type EdgeGraph struct {
	edges []*Edge
}

// Edge - edge of EdgeGraph
type Edge struct {
	VertexFrom *Vertex
	VertexTo   *Vertex
	Weight     int64
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

// GetEdges - get all edges
func (g *EdgeGraph) GetEdges() []*Edge {
	return g.edges
}

// GetVertexByValue - get vertex by value
func (g *EdgeGraph) GetVertexByValue(value int64) *Vertex {
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

// GetEdge - get edge by values of edges
func (g *EdgeGraph) GetEdge(vertexFromValue int64, vertexToValue int64) (*Edge, error) {
	for _, edge := range g.edges {
		if edge.VertexFrom.Value == vertexFromValue && edge.VertexTo.Value == vertexToValue {
			return edge, nil
		}
	}
	return nil, fmt.Errorf("unable to find edge with '%d' and '%d'  vertices", vertexFromValue, vertexToValue)
}

func (g *EdgeGraph) addEdge(vertexFrom *Vertex, vertexTo *Vertex, weight int64) {
	for _, edge := range g.edges {
		if edge.VertexFrom == vertexFrom && edge.VertexTo == vertexTo {
			return
		}
	}
	g.edges = append(g.edges, &Edge{
		VertexFrom: vertexFrom,
		VertexTo:   vertexTo,
		Weight:     weight,
	})
}
