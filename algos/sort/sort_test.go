package sort_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parallelKiller/fundamentals/algos/sort"
)

var _ = Describe("Sort", func() {

	Describe("QuickSort", func() {
		var sorted []int
		var unsorted []int
		BeforeEach(func() {
			unsorted = []int{0, 0, 39, -421, 01, 9000, 44, -44}
			sorted = []int{-421, -44, 0, 0, 1, 39, 44, 9000}
		})

		It("sorts the slice", func() {
			Expect(sort.QuickSort(unsorted)).To(Equal(sorted))
		})

	})
})
