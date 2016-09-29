package pqueue_test

import (
	"github.com/SpectralHiss/fundamentals/adt/pqueue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
	we needed a min-queue for dijikstra will later backfill normal tests
	queue is storing an id for each entry
*/

var _ = Describe("minimum Priority queue ", func() {

	var (
		pq pqueue.MinPQueue

		testEntries = []int{4, 2, 53, 1000, 42, -1, -3000}
	)

	BeforeEach(func() {
		pq = pqueue.NewMinQueue()

		for i := 0; i < len(testEntries); i++ {
			pq.Insert(testEntries[i], i)
		}

	})

	It("dequeues lowest priority element first", func() {
		Expect(pq.ExtractMin()).To(Equal(6))
		Expect(pq.ExtractMin()).To(Equal(5))
		Expect(pq.ExtractMin()).To(Equal(1))
	})

	It("decreases key properly", func() {
		pq.DecreaseK(3, -3421)
		Expect(pq.ExtractMin()).To(Equal(3))
		Expect(pq.ExtractMin()).To(Equal(6))
	})

})
