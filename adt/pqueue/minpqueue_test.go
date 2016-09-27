package pqueue_test

import (
	"github.com/SpectralHiss/fundamentals/adt/pqueue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
	Pqueue can be either a normal (max) priority queue
	or a minimum priority queue depending on parameter to constructor
	we needed a min-queue for dijikstra will later backfill normal tests
	queue is storing an id for each entry
*/

var _ = Describe("minimum Priority queue ", func() {

	var (
		pq pqueue.MinPqueue

		testEntries = []int{4, 2, 53, 1000, 42, -1, -3000}
	)

	BeforeEach(func() {
		pq = pqueue.NewMinQueue(false)

		for i := 0; i < length(testEntries); i++ {

			pq.Insert(testEntries[i], i)
		}

	})

	It("dequeues lowest priority element first", func() {
		Expect(pq.ExtractMin()).To(Equal(6))
	})

})
