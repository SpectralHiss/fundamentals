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
	stack.content = prepend(stack.content, item)
	stack.size++
	return nil

}

func (stack *SimpleStack) Pop() (interface{}, error) {
	if stack.size == 0 {
		return 0, fmt.Errorf("empty stack popped")
	}

	return stack.content[0], nil
}

func prepend(slice []interface{}, elem interface{}) []interface{} {
	tempSlice := make([]interface{}, cap(slice))
	tempSlice[0] = elem

	for index, elem := range slice[0 : len(slice)-1] {
		tempSlice[index+1] = elem
	}
	return tempSlice
}
