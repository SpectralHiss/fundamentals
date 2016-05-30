package set_test

import (
	"github.com/parallelKiller/fundamentals/adt/set"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	Describe("Bit set, limited to mapping N 0 -> 64, but efficient..", func() {
		var bitset *set.BitVector
		BeforeEach(func() {
			bitset = set.New()
		})

		Context("when an element is inserted to the set", func() {
			BeforeEach(func() {
				bitset.Add(uint(10))
			})
			It("is a member of the set", func() {
				Expect(bitset.Member(uint(10))).To(Equal(true))
			})
		})
	})

})
