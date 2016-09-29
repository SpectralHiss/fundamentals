package heap_test

import (
	"github.com/SpectralHiss/fundamentals/datastructures/heap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Min Heap", func() {
	var minHeap heap.MinHeap

	BeforeEach(func() {
		minHeap = heap.NewMinHeap()

		testEntries := []int{4, -2, 13, 100, 423, -1, -3000}

		for index := 0; index < len(testEntries); index++ {
			minHeap.Insert(heap.Entry{W: testEntries[index], Id: index})
		}
	})

	It("allows peeking to the minimum weight element", func() {
		Expect(minHeap.Min().Id).To(Equal(6))

	})

	It("keeps the min Heap constraint", func() {
		Expect(MinHeapConstraintRec(1, minHeap.(*heap.ArrMinHeap))).To(BeTrue())
	})

	It("removes the minimum element on extract min, keeping the constraint", func() {
		//	fmt.Printf("%#v", minHeap)
		minElem := minHeap.ExtractMin()
		Expect(minElem.Id).To(Equal(6))
		Expect(minHeap.Min().Id).NotTo(Equal(6))
		Expect(MinHeapConstraintRec(1, minHeap.(*heap.ArrMinHeap))).To(BeTrue())
	})

})

func MinHeapConstraintRec(i int, minh *heap.ArrMinHeap) bool {
	Heap := *minh

	if i >= len(Heap)/2 {
		return true
	}

	if (Heap)[2*i].W < Heap[i].W || Heap[2*i+1].W < Heap[i].W {
		return false
	} else {
		return MinHeapConstraintRec(2*i, minh) && MinHeapConstraintRec(2*i+1, minh)
	}
}
