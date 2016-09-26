package bst_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/SpectralHiss/fundamentals/algos/sort"
	"github.com/SpectralHiss/fundamentals/datastructures/bst"
	"github.com/SpectralHiss/fundamentals/datastructures/tree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// this tree does not balance and does not permit dupes.
var (
	cleverTree bst.BST
	head       tree.Node = tree.Node{K: 10, V: "Head"}
	chiffre    tree.Node = tree.Node{K: 22, V: "le chiffre"}
	colours    tree.Node = tree.Node{K: 19, V: "colours"}
	weights    []int
)

var _ = Describe("Bst", func() {

	BeforeEach(func() {
		cleverTree = bst.New(head)
		cleverTree.Add(chiffre)
		cleverTree.Add(colours)

		addRandBunch(cleverTree)
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
		weights = getWeights(cleverTree.Flatten())
		Expect(weights).To(Equal(sort.QuickSort(weights)))
	})

	Context("deletion", func() {
		var Peugeot, Dalmatiens, Judas, Glutton, Gluten tree.Node

		itContains := func(weights []int, nodes []tree.Node) {
			for _, elem := range nodes {
				Expect(weights).To(ContainElement(elem.K))
			}
		}

		BeforeEach(func() {
			/* the idea is to build attached to the right of the big
			tree something that looks like :
			   		RestOfTree ( K  < 100)...
								\
								 \
								Peugeot
								/	\
							   /	 \
						Dalmatiens	  Null
						/	\
					   /	 \
					Null		Judas
								/	\
							   /	 \
						Glutton		  Gluten
						/	\		/		\
					   Null	Null	Null	Null

			 this gives us all edge cases we need
			 numerology is stupid , lord have mercy*/

			Peugeot = tree.Node{K: 105, V: "Peugeot"}
			Dalmatiens = tree.Node{K: 101, V: "Dalmatiens"}
			Judas = tree.Node{K: 103, V: "Judas"}
			Glutton = tree.Node{K: 102, V: "Glutton"}
			Gluten = tree.Node{K: 104, V: "Gluten-free!"}

			cleverTree.Add(Peugeot)
			cleverTree.Add(Dalmatiens)
			cleverTree.Add(Judas)
			cleverTree.Add(Glutton)
			cleverTree.Add(Gluten)
		})

		Context("when the element deleted is a leaf node", func() {
			BeforeEach(func() {
				cleverTree.Remove(Gluten)
				weights = getWeights(cleverTree.Flatten())
			})

			It("deletes element properly", func() {
				Expect(weights).NotTo(ContainElement(Gluten.K))

				itContains(weights, []tree.Node{
					Peugeot, Judas, Glutton, Dalmatiens,
				})

				Expect(weights).To(Equal(sort.QuickSort(weights)))
			})
		})

		Context("when the element deleted is the head node", func() {

			BeforeEach(func() {
				cleverTree.Remove(head)
				weights = getWeights(cleverTree.Flatten())
			})

			It("deletes element properly", func() {
				Expect(weights).NotTo(ContainElement(head))

				itContains(weights, []tree.Node{
					Peugeot, Judas, Glutton, Dalmatiens, Gluten,
				})

				Expect(weights).To(Equal(sort.QuickSort(weights)))
			})
		})

		Context("when the element deleted has no right subtree", func() {
			BeforeEach(func() {
				cleverTree.Remove(Peugeot)
				weights = getWeights(cleverTree.Flatten())
			})

			It("deletes the element properly", func() {
				Expect(weights).NotTo(ContainElement(Peugeot.K))

				itContains(weights, []tree.Node{
					Dalmatiens, Judas, Glutton, Gluten,
				})

				Expect(weights).To(Equal(sort.QuickSort(weights)))
			})

		})

		Context("when the element deleted has no left subtree", func() {
			BeforeEach(func() {
				cleverTree.Remove(Dalmatiens)
				weights = getWeights(cleverTree.Flatten())
			})

			It("deletes element properly", func() {
				Expect(weights).NotTo(ContainElement(Dalmatiens.K))

				itContains(weights, []tree.Node{
					Peugeot, Judas, Glutton, Gluten,
				})

				Expect(weights).To(Equal(sort.QuickSort(weights)))
			})
		})

		Context("when the element deleted has left and right subtree", func() {
			BeforeEach(func() {
				cleverTree.Remove(Judas)
				weights = getWeights(cleverTree.Flatten())
			})

			It("deletes element properly", func() {
				Expect(weights).NotTo(ContainElement(Judas.K))

				itContains(weights, []tree.Node{
					Peugeot, Dalmatiens, Glutton, Gluten,
				})

				Expect(weights).To(Equal(sort.QuickSort(weights)))
			})
		})
	})
})

func addRandBunch(cleverTree bst.BST) {
	limit := 10
	randGen := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	sieve := make([]bool, 100)

	for {
		if limit <= 0 {
			break
		}

		var key int
		// this is to avoid dupes
		for {
			key = int(randGen.Float64() * 100)

			if !sieve[key] {
				sieve[key] = true
				break
			}
		}

		cleverTree.Add(tree.Node{K: key, V: fmt.Sprintf("i:%d", limit)})
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
