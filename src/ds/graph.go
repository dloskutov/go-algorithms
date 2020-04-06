package ds

type vertex struct {
	value            int64
	adjacentVertices []*vertex
}

// Graph - undirected graph
type Graph struct {
	vertices []*vertex
}

// NewGraph - create graph based on edges
func NewGraph(edges [][]int64) *Graph {
	graph := &Graph{
		vertices: make([]*vertex, 0),
	}

	for _, edge := range edges {
		leftVertex := graph.getVertexByValue(edge[0])
		rightVertex := graph.getVertexByValue(edge[1])

		if leftVertex == nil {
			leftVertex = &vertex{
				value:            edge[0],
				adjacentVertices: make([]*vertex, 0),
			}
			graph.vertices = append(graph.vertices, leftVertex)
		}
		if rightVertex == nil {
			rightVertex = &vertex{
				value:            edge[1],
				adjacentVertices: make([]*vertex, 0),
			}
			graph.vertices = append(graph.vertices, rightVertex)
		}

		leftVertex.addAdjacentVertex(rightVertex)
		rightVertex.addAdjacentVertex(leftVertex)
	}
	return graph
}

// HasVertexWithValue - check if graph has vertex with specific value
func (g *Graph) HasVertexWithValue(value int64) bool {
	for _, vertex := range g.vertices {
		if vertex.value == value {
			return true
		}
	}
	return false
}

func (g *Graph) getVertexByValue(value int64) *vertex {
	for _, vertex := range g.vertices {
		if vertex.value == value {
			return vertex
		}
	}
	return nil
}

func (v *vertex) addAdjacentVertex(vertexToAdd *vertex) {
	hasVertex := false
	for _, vertex := range v.adjacentVertices {
		if vertex == vertexToAdd {
			hasVertex = true
			break
		}
	}
	if !hasVertex {
		v.adjacentVertices = append(v.adjacentVertices, vertexToAdd)
	}
}
