package linkedlist_test

import (
	llist "github.com/parallelKiller/fundamentals/datastructures/linkedlist"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LinkedList", func() {
	var list llist.LinkedList
	BeforeEach(func() {
		list = llist.New()
	})

	Describe("single elem add, first and last", func() {

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
	})

	Describe("linkedlist many elements, addition and search", func() {
		Context("when two elements are added to the list", func() {

			BeforeEach(func() {
				list.Add(22)
				list.Add(22)
				list.Add(77)
			})

			It("returns the elements consecutively in a FIFO = LILO fashion", func() {
				Expect(list.First()).To(Equal(22))
				Expect(list.Next()).To(Equal(22))
				Expect(list.Last()).To(Equal(77))
			})
		})

		It("returns indeces of elements", func() {
			list.AddMany(422, 99, 135, 442, 135, 1, 1)
			Expect(list.First()).To(Equal(422))
			Expect(list.Last()).To(Equal(1))

			Expect(list.GetIndeces(1)).To(Equal([]int{5, 6}))
		})
	})

	Describe("linkedlist element deletion", func() {
		BeforeEach(func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
		})

		Context("when the first elemen is deleted", func() {
			BeforeEach(func() {
				list.DeleteAt(0)
			})
			It("should be dropped from the list", func() {
				Expect(list.First()).To(Equal(2))
				Expect(list.Last()).To(Equal(3))
			})
		})

		Context("when the second element is deleted", func() {
			BeforeEach(func() {
				list.DeleteAt(1)
			})

			It("should be dropped from the list", func() {
				Expect(list.First()).To(Equal(1))
				Expect(list.Next()).To(Equal(3))
			})
		})

		Context("when the third element is deleted", func() {
			BeforeEach(func() {
				list.DeleteAt(2)
			})

			It("should be dropped from the list", func() {
				Expect(list.First()).To(Equal(1))
				Expect(list.Last()).To(Equal(2))
			})
		})

	})

})
