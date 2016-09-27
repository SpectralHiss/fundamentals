package pqueue

import "github.com/SpectralHiss/datastructure/heap"

type Entry struct {
	Priority int
}

type MinPQueue interface {
	Insert(priority int, id int ) 
	ExtractMin() int 
	DecreaseK(index int)
}

func NewMinQueue() {
	heap := (maxQueue)? heap.MinHeap() : heap.MaxHeap()

}


func (pq *Pqueue) Insert(priority int, id int) {

}
