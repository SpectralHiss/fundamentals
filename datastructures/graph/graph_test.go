package graph_test

import (
	"github.com/SpectralHiss/fundamentals/datastructures/graph"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Graph", func() {
	var g graph.Graph

	BeforeEach(func() {
		g = graph.New([][]int{{0, 1, 0, 1, 0, 0},
			{0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1, 7},
			{0, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 1}})

	})

	It("returns the correct adjacency list for a particular node (0-indexed)", func() {
		adj := g.AdjList(2)
		Expect(adj).To(Equal([]graph.Edge{
			{EndNode: 4, Distance: 1}, {EndNode: 5, Distance: 7},
		}))
	})
})
