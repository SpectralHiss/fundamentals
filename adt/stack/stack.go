package stack

import "fmt"

type SimpleStack struct {
	content []interface{}
	size    int
}

func New(capacity int) SimpleStack {
	return SimpleStack{
		content: make([]interface{}, capacity),
	}
}

func (stack *SimpleStack) Push(item interface{}) error {
	if stack.size == len(stack.content) {
		return fmt.Errorf("cannot push: stack is full")
	}
	stack.content = prepend(stack.content, item)
	stack.size++
	return nil
}

func (stack *SimpleStack) Pop() (interface{}, error) {
	if stack.size == 0 {
		return 0, fmt.Errorf("empty stack popped")
	}
	ret := stack.content[0]
	stack.content = stack.content[1:]
	return ret, nil
}

func prepend(slice []interface{}, elem interface{}) []interface{} {

	if cap(slice) == 0 {
		return []interface{}{elem}
	}
	tempSlice := make([]interface{}, cap(slice))
	tempSlice[0] = elem

	for index, elem := range slice[0 : len(slice)-1] {
		tempSlice[index+1] = elem
	}

	return tempSlice
}
