package ds

// Graph - undirected graph
type Graph struct {
	vertices []*Vertex
}

// NewGraph - create graph based on edges
func NewGraph(edges [][]int64) *Graph {
	graph := &Graph{
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
		rightVertex.addAdjacentVertex(leftVertex)
	}
	return graph
}

// GetVertices - get all graph vertices
func (g *Graph) GetVertices() []*Vertex {
	return g.vertices
}

// HasVertexWithValue - check if graph has vertex with specific value
func (g *Graph) HasVertexWithValue(value int64) bool {
	for _, vertex := range g.vertices {
		if vertex.Value == value {
			return true
		}
	}
	return false
}

// GetVertexByValue - get vertex by value
func (g *Graph) GetVertexByValue(value int64) *Vertex {
	for _, vertex := range g.vertices {
		if vertex.Value == value {
			return vertex
		}
	}
	return nil
}

// Vertex - graph vertex
type Vertex struct {
	Value            int64
	AdjacentVertices []*Vertex
}

func (v *Vertex) addAdjacentVertex(vertexToAdd *Vertex) {
	hasVertex := false
	for index := range v.AdjacentVertices {
		if v.AdjacentVertices[index] == vertexToAdd {
			hasVertex = true
			break
		}
	}
	if !hasVertex {
		v.AdjacentVertices = append(v.AdjacentVertices, vertexToAdd)
	}
}
