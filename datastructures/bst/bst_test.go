package bst_test

import (
	"fmt"
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parallelKiller/fundamentals/algos/sort"
	"github.com/parallelKiller/fundamentals/datastructures/bst"
	"github.com/parallelKiller/fundamentals/datastructures/tree"
)

var _ = Describe("Bst", func() {
	// this tree does not balance and does not permit dupes.
	var (
		cleverTree bst.BST
		head       tree.Node = tree.Node{K: 10, V: "Head"}
		chiffre    tree.Node = tree.Node{K: 22, V: "le chiffre"}
		colours    tree.Node = tree.Node{K: 19, V: "colours"}
		weights    []int
	)

	BeforeEach(func() {
		cleverTree = bst.New(head)
		cleverTree.Add(chiffre)
		cleverTree.Add(colours)

		addRandBunch(cleverTree)

		weights = getWeights(cleverTree.Flatten())
	})

	It("allows adding elements", func() {
		Expect(cleverTree.Head()).To(Equal(head))
		Expect(cleverTree.Right().Head()).To(Equal(chiffre))
		Expect(cleverTree.Right().Left().Head()).To(Equal(colours))
	})

	It("is searchable ", func() {
		Expect(cleverTree.Search(chiffre)).To(Equal(cleverTree.Right()))
	})

	It("maintains ordering (left node < top node < right node) ", func() {
		Expect(weights).To(Equal(sort.QuickSort(weights)))
	})

	It("allows deleting elements", func() {
		cleverTree.Remove(chiffre)
		Expect(cleverTree.Head()).To(Equal(head))
		Expect(cleverTree.Right().Head()).To(Equal(colours))
	})

})

func addRandBunch(cleverTree bst.BST) {
	limit := 10
	for {
		if limit <= 0 {
			break
		}
		cleverTree.Add(tree.Node{K: rand.Int(), V: fmt.Sprintf("i:%d", limit)})
		limit--
	}
}

func getWeights(splat []tree.Node) []int {

	// sadly no "map" in go
	weights := []int{}
	for _, elem := range splat {
		weights = append(weights, elem.K)
	}
	return weights
}
