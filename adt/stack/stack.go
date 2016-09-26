package stack

import (
	"fmt"
	"github.com/SpectralHiss/fundamentals/helpers"
)

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
	stack.content = helpers.PrependNotGrow(stack.content, item)
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
