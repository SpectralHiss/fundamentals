package stack_test

import (
	"github.com/spectreOfAbsorbance/ADTalGOstudy/adt/stack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack datastructure tests", func() {


	Context("when the stack size is fixed", func(){
		var simplestack stack.SimpleStack

				simplestack = stack.New(10)

			Context("when the stack already contains items", func(){
				BeforeEach(func(){
					simplestack.Push(360)
				})

				It("should pop the last item pushed first", func (){
					elem , err := simplestack.Pop()
					Expect(err).To(BeNil())
					Expect(elem).To(Equal(360))
				})
			})

			Context("when the stack is empty",func(){

				BeforeEach(func(){
					simplestack = stack.New(10)
				})

				Context("and the stack is popped", func(){
					var err error
					BeforeEach(func(){
					 _, err = simplestack.Pop()
				 })

					It("returns an error",func(){
						Expect(err).NotTo(BeNil())
					})

					It("has a helplful error message", func(){
					})
				})
			})


			})


		})
