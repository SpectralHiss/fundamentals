package bst_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parallelKiller/fundamentals/datastructures/bst"
	"github.com/parallelKiller/fundamentals/datastructures/tree"
)

var _ = Describe("Bst", func() {

	var cleverTree bst.BST

	It("allows adding elements", func() {
		cleverTree = bst.New(tree.Node{K: 10, V: "Head"})
		cleverTree.Add(tree.Node{K: 22, V: "le chiffre"})

		Expect(cleverTree.Right().Head().K).To(Equal(22))
		Expect(cleverTree.Right().Head().V).To(Equal("le chiffre"))
	})

	It("maintains ordering (left node < top node < right node) ", func() {
		// DFS should give gradually >= elements..

	})
})
