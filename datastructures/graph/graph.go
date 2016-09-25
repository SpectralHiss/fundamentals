package graph

// let's start with simple DAG

type Edge struct {
	EndNode  int // number of vertex
	Distance int
}

type DWGraph struct {
	matrix [][]int
}

type Graph interface {
	AddEdge(int, Edge)
	AdjList(node int) []Edge
	Matrix() [][]int
}

func New(InitialMatrix [][]int) Graph {
	return DWGraph{
		matrix: InitialMatrix,
	}
}

func (graph DWGraph) AddEdge(node int, edge Edge) {
	graph.matrix[node][edge.EndNode] = edge.Distance
}

func (graph DWGraph) Matrix() [][]int {
	return graph.matrix
}

func (graph DWGraph) AdjList(node int) []Edge {
	var adj []Edge
	for node, distance := range graph.matrix[node] {
		if distance != 0 {
			adj = append(adj, Edge{EndNode: node, Distance: distance})
		}
	}
	return adj
}
