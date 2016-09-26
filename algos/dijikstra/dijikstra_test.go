package dijikstra_test

import (
	"github.com/SpectralHiss/fundamentals/algos/dijikstra"
	"github.com/SpectralHiss/fundamentals/datastructures/graph"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = PDescribe("Dijikstra", func() {
	var directedG graph.Graph

	BeforeEach(func() {
		directedG = graph.New([][]int{
			{0, 1, 5, 0, 0, 0, 0},
			{0, 0, 0, 7, 0, 0, 18},
			{0, 0, 0, 12, 0, 0, 0},
			{0, 0, 0, 0, 2, 3, 0},
			{0, 0, 0, 0, 0, 5, 0},
			{0, 0, 0, 0, 0, 0, 2},
			{0, 0, 0, 0, 0, 0, 0}})
	})

	It("computes the shortest path between node 0 and 6", func() {
		path := dijikstra.ShortestPath(directedG, 0)

		Expect(path).To(Equal(dijikstra.Path{0, 1, 3, 5, 6}))
	})

})
