package queue

import llist "github.com/parallelKiller/fundamentals/datastructures/linkedlist"

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Peek() interface{}
}

type Simplequeue struct {
	lnklist llist.LinkedList
}

func New() Simplequeue {
	return Simplequeue{}
}

func (s *Simplequeue) Enqueue(elem interface{}) {
	s.lnklist.Add(elem)
}

// TODO: actually dequeue, pass pending test
func (s *Simplequeue) Dequeue() interface{} {
	return s.lnklist.First()
}

func (s *Simplequeue) Peek() interface{} {
	return s.lnklist.First()
}
