package stack_test

import (
	"github.com/parallelKiller/fundamentals/adt/stack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack datastructure tests", func() {

	Context("when the stack size is fixed", func() {
		var simplestack stack.SimpleStack

		simplestack = stack.New(10)

		Context("when the stack contains an item", func() {
			BeforeEach(func() {
				simplestack.Push(360)
			})

			It("should pop the last item pushed first", func() {
				elem, err := simplestack.Pop()
				Expect(err).To(BeNil())
				Expect(elem).To(Equal(360))
			})
		})

		Context("when the stack contains two items", func() {
			BeforeEach(func() {
				simplestack.Push(1080)
				simplestack.Push(7)
			})

			It("should pop first item second, in a FILO = LIFO fashion", func() {
				elem, err := simplestack.Pop()
				Expect(err).To(BeNil())
				Expect(elem).To(Equal(7))
				elem, err = simplestack.Pop()
				Expect(err).NotTo(HaveOccurred())
				Expect(elem).To(Equal(1080))

			})
		})

		Context("when the stack is empty", func() {

			BeforeEach(func() {
				simplestack = stack.New(10)
			})

			Context("and the stack is popped", func() {
				var err error
				BeforeEach(func() {
					_, err = simplestack.Pop()
				})

				It("returns an error", func() {
					Expect(err).NotTo(BeNil())
				})

				It("has a helplful error message", func() {
					Expect(err).To(MatchError("empty stack popped"))
				})
			})
		})

		Context("when the stack is full", func() {

			BeforeEach(func() {
				for i := 0; i <= 10; i++ {
					simplestack.Push(i)
				}
			})

			Context("and an item is pushed", func() {
				var err error
				BeforeEach(func() {
					err = simplestack.Push(11)
				})

				It("returns an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})

	})

})
