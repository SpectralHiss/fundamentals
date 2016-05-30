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
	return Simplequeue{
		lnklist: llist.New(),
	}
}

func (s *Simplequeue) Enqueue(elem interface{}) {
	s.lnklist.Add(elem)
}

func (s *Simplequeue) Dequeue() interface{} {
	l := s.lnklist
	ret := l.First()
	l.DeleteAt(0)
	return ret
}

func (s *Simplequeue) Peek() interface{} {
	return s.lnklist.First()
}
