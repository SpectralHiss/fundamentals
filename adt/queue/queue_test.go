package queue_test

import (
	queue "github.com/SpectralHiss/fundamentals/adt/queue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queue", func() {
	Context("When three elements are queued", func() {

		myqueue := queue.New()
		myqueue.Enqueue(1)
		myqueue.Enqueue(2)
		myqueue.Enqueue(3)

		It("should dequeue the first element in the queue", func() {
			Expect(myqueue.Dequeue()).To(Equal(1))
			Expect(myqueue.Dequeue()).To(Equal(2))
			Expect(myqueue.Peek()).To(Equal(3))
		})

	})
})
