package pqueue

import "github.com/SpectralHiss/fundamentals/datastructures/heap"

type MinPQueue interface {
	Insert(priority int, id int)
	ExtractMin() (id int)
	DecreaseK(id int, newK int)
}

type MinPQ struct {
	heap.MinHeap
}

func NewMinQueue() MinPQueue {
	heap := heap.NewMinHeap()
	return &MinPQ{heap}
}

func (pq *MinPQ) ExtractMin() int {
	return pq.MinHeap.ExtractMin().Id
}

// this could have been implemented inside minheap.. not sure if it matters!
func (pq *MinPQ) DecreaseK(index int, newK int) {
	pos, elem := pq.Scan(index)
	pq.Remove(pos)
	elem.W = newK
	pq.MinHeap.Insert(elem)

}

func (pq *MinPQ) Insert(priority int, id int) {
	pq.MinHeap.Insert(heap.Entry{W: priority, Id: id})
}
