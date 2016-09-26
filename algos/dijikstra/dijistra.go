package dijikstra

import (
	"github.com/SpectralHiss/fundamentals/datastructures/graph"
	"github.com/SpectralHiss/fundamentals/helpers"
)

type Path []int

type NodeDesc struct {
	Dist int
	Pred int
}

type NodeDescs []NodeDesc

func (descs NodeDescs) ToPath(endNode int) {
	path := Path{}

	var curNode int
	for curNode = endNode; curNode != -1; curNode = descs[curNode].Pred {
		path = helpers.PrependAndGrow(path, curNode)
	}
}

func initialise() {

}

func ShortestPath(graph graph.Graph, start int) Path {
	NodeDescs := make([]int, len(graph.AdjList(0)))

	return NodeDescs
}
