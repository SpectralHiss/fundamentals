package queue

import llist "github.com/spectreOfAbsorbance/ADTalGOstudy/DataStructures/LinkedList"


type Queue interface{
	Enqueue(interface{})
	Dequeue() interface{}
}


type Simplequeue struct{
	lnklist llist.LinkedList
}


func New() Simplequeue {
	return Simplequeue{}
}

func (s *Simplequeue) Enqueue (elem interface{}){
	s.lnklist.Add(elem)
}

func (s *Simplequeue) Dequeue () interface{} {

	return s.lnklist.First()
}

