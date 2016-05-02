package stack_test

import (
	"github.com/spectreOfAbsorbance/fundamentals/adt/helpers"
	"github.com/spectreOfAbsorbance/fundamentals/adt/stack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("checking that our stack can handle combinations of types", func() {
	manyTypes := helpers.ManyTypes()

	var simplestack stack.SimpleStack
	BeforeEach(func() {
		simplestack = stack.New(10)
	})
	Context("when an item of any type is pushed to the stack", func() {
		BeforeEach(func() {
			for _, elem := range manyTypes {
				simplestack.Push(elem)
			}

		})

		It("returns the same items when popped", func() {
			var topOfStack interface{}
			for j, _ := range manyTypes {
				var err error
				topOfStack, err = simplestack.Pop()

				Expect(err).ToNot(HaveOccurred())

				Expect(topOfStack).To(Equal(manyTypes[len(manyTypes)-(j+1)]))
			}

		})

	})
})
