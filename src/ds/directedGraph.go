package ds

// DirectedGraph - directed graph
type DirectedGraph struct {
	vertices []*Vertex
}

// NewDirectedGraph - create directed graph based on edges
func NewDirectedGraph(edges [][]int64) *DirectedGraph {
	graph := &DirectedGraph{
		vertices: make([]*Vertex, 0),
	}

	for _, edge := range edges {
		leftVertex := graph.GetVertexByValue(edge[0])
		rightVertex := graph.GetVertexByValue(edge[1])

		if leftVertex == nil {
			leftVertex = &Vertex{
				Value:            edge[0],
				AdjacentVertices: make([]*Vertex, 0),
			}
			graph.vertices = append(graph.vertices, leftVertex)
		}
		if rightVertex == nil {
			rightVertex = &Vertex{
				Value:            edge[1],
				AdjacentVertices: make([]*Vertex, 0),
			}
			graph.vertices = append(graph.vertices, rightVertex)
		}

		leftVertex.addAdjacentVertex(rightVertex)
	}
	return graph
}

// GetVertices - get all vertices
func (g *DirectedGraph) GetVertices() []*Vertex {
	return g.vertices
}

// HasVertexWithValue - check if graph has vertex with specific value
func (g *DirectedGraph) HasVertexWithValue(value int64) bool {
	for _, vertex := range g.vertices {
		if vertex.Value == value {
			return true
		}
	}
	return false
}

// GetVertexByValue - get vertex by value
func (g *DirectedGraph) GetVertexByValue(value int64) *Vertex {
	for _, vertex := range g.vertices {
		if vertex.Value == value {
			return vertex
		}
	}
	return nil
}

// HasEdge - check if edge exists
func (g *DirectedGraph) HasEdge(from int64, to int64) bool {
	fromVertex := g.GetVertexByValue(from)
	if fromVertex == nil {
		return false
	}
	for _, toVertex := range fromVertex.AdjacentVertices {
		if toVertex.Value == to {
			return true
		}
	}

	return false
}
