package linkedlist

// import "fmt"

type LinkedListNode struct {
	next *LinkedListNode
	curr *interface{}
}

type LinkedList interface {
	Add(interface{})
	AddMany(...interface{})
	DeleteAt(int)
	GetIndeces(interface{}) []int
	First() interface{}
	Next() interface{}
	Last() interface{}
}

func New() *LinkedListNode {
	return &LinkedListNode{
		next: nil,
		curr: nil,
	}
}

func (list *LinkedListNode) AddMany(elems ...interface{}) {
	if len(elems) == 1 {
		list.Add(elems[0])
	} else {
		list.Add(elems[0])
		list.AddMany(elems[1:len(elems)]...)
	}
}

func (list *LinkedListNode) Add(elem interface{}) {
	newlist := LinkedListNode{
		next: nil,
		curr: &elem,
	}

	if list.curr == nil && list.next == nil {
		*list = newlist
	} else {
		list.lastlist().next = &newlist
	}
}

func (list *LinkedListNode) DeleteAt(index int) {
	if index == 0 {
		*list = *list.next
	} else {
		cursor := 0
		currList := list
		for {
			if cursor+1 == index {
				// insertion poin t
				curr := currList.next

				*currList.next = *curr.next
				return
			}
			cursor += 1
			currList = currList.next
		}

	}

}

func (list *LinkedListNode) GetIndeces(elem interface{}) []int {
	var positions []int
	count := 0

	currList := list

	for {
		if *currList.curr == elem {
			positions = append(positions, count)
		}

		if currList.next == nil {
			break
		}
		currList = currList.next
		count++
	}
	return positions
}

func (list *LinkedListNode) lastlist() *LinkedListNode {
	currList := list
	for {
		if currList.next == nil {
			return currList
		}
		currList = currList.next
	}
}

func (list *LinkedListNode) Last() interface{} {
	return *(list.lastlist().curr)
}

func (list *LinkedListNode) First() interface{} {
	return *list.curr
}

func (list *LinkedListNode) Next() interface{} {
	return *list.next.curr
}

func (list *LinkedListNode) NextElem() interface{} {
	return *((*list.next).curr)
}
