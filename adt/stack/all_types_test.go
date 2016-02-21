package stack_test

import (
	"github.com/spectreOfAbsorbance/ADTalGOstudy/adt/stack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func nonTrivialMap() map[int]map[interface{}]interface{} {
	return map[int]map[interface{}]interface{}{
		3: map[interface{}]interface{}{
			"aluminium": "neuro-toxic poison",
			33:          "spies",
		},
	}
}

func channel() chan bool {
	c := make(chan bool, 10)
	return c
}
func array() [4]int {
	return [...]int{1, 2, 3, 4}
}

var _ = Describe("checking that our stack can handle combinations of types", func() {
	manyTypes := []interface{}{
		0, "1", complex(10, 5), float64(19.44443), nonTrivialMap(), channel(), array(),
	}
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
