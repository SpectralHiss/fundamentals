package linkedlist_test

import (
	"fmt"

	llist "github.com/spectreOfAbsorbance/fundamentals/datastructures/linkedlist"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LinkedList", func() {
	var list *llist.LinkedList
	BeforeEach(func() {
		list = llist.New()

	})

	Context("when an element is added to the list", func() {
		BeforeEach(func() {
			list.Add(2)
		})
		It("is first in the list", func() {
			Expect(list.First()).To(Equal(2))
		})

		It("is last in the list", func() {
			Expect(list.Last()).To(Equal(2))
		})
	})

	Context("when two elements are added to the list", func() {

		BeforeEach(func() {
			list.Add(22)
			list.Add(77)
		})

		It("returns the elements consecutively in a FIFO = LILO fashion", func() {
			Expect(list.First()).To(Equal(22))
			Expect(list.Next()).To(Equal(77))
		})
	})

	PIt("returns indeces of elements", func() {
		list.Add(422)
		list.Add(99)
		list.Add(135)
		list.Add(442)
		list.Add(135)
		list.Add(1)
		list.Add(1)
		fmt.Printf("%+v", list)
		Expect(list.First()).To(Equal(422))
		Expect(list.Last()).To(Equal(42))
		// Expect(list.getIndeces(1)).To(Equal([]int{5,6})
	})

})
