package linkedlist

type LinkedList struct{
	next *LinkedList
	curr *interface{}
}


func New() LinkedList {
	return LinkedList{
		next: nil,
		curr: nil,
	}
}

func (list LinkedList) Add(elem interface{}){
	newlist := LinkedList {
		next:nil,
		curr: &elem,
	}
	
	if list.Next() == nil {
		list = newlist
	} else {
	*list.Last().next = newlist
	}

}



func (list LinkedList) Last() *LinkedList{
	for {
	if list := list.Next() ; list.Next() == nil {

		return list
		}
	}
}

func (list LinkedList) First() interface{}{
	return *list.curr
}

func (list LinkedList) Next() *LinkedList {
	return list.next
}

func ( list LinkedList) NextElem() interface{}{
	return *((*list.next).curr)
}
