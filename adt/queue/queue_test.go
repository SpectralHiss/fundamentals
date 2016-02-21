package queue_test

import (
	queue "github.com/spectreOfAbsorbance/ADTalGOstudy/adt/queue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queue", func() {
	Context("When two elements are queued",func(){
		
		myqueue := queue.New()
		myqueue.Enqueue(1)
		myqueue.Enqueue(2)
		
		It("should dequeue the first element in the queue" , func(){
			Expect(myqueue.Dequeue()).To(Equal(1))	
		})
	})
})
