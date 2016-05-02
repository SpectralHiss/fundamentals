package linkedlist

type LinkedList struct {
	next *LinkedList
	curr *interface{}
}

func New() *LinkedList {
	return &LinkedList{
		next: nil,
		curr: nil,
	}
}

// func (list *LinkedList) AddMany(elems ...interface{}) {
// 	if len(elems) == 1 {
// 		list.Add(elems)
// 	} else {
// 		list.Add(elems[0])
// 		fmt.Printf("%#v", elems[1:])
// 		list.AddMany(elems[1:len(elems)]...)

// 	}
// }

func (list *LinkedList) Add(elem interface{}) {
	newlist := LinkedList{
		next: nil,
		curr: &elem,
	}

	if list.curr == nil && list.next == nil {
		*list = newlist
	} else {
		list.lastlist().next = &newlist
	}
}

//TODO: finish!
func (list *LinkedList) getIndeces(elem interface{}) []int {
	var positions []int
	count := 0

	if list.curr == elem {
		positions = append(positions, count)
	}

	return positions
}

func (list *LinkedList) lastlist() *LinkedList {
	currList := list.next
	for {
		if currList == nil {
			return list
		} else if currList.next == nil {
			return currList
		}
		currList = list.next
	}
}
func (list *LinkedList) Last() interface{} {
	return *(list.lastlist().curr)
}

func (list *LinkedList) First() interface{} {
	return *list.curr
}

func (list *LinkedList) Next() interface{} {
	return *list.next.curr
}

func (list *LinkedList) NextElem() interface{} {
	return *((*list.next).curr)
}
