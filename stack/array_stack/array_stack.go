package array_stack

import "errors"

type ArrayStack struct {
	elements []interface{}
}

func New() *ArrayStack {
	return &ArrayStack{}
}

func (stack *ArrayStack) Empty() bool {
	return len(stack.elements) == 0
}

func (stack *ArrayStack) Top() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("this stack is empty, can't top the element")
	}
	return stack.elements[len(stack.elements)-1], nil
}

func (stack *ArrayStack) Pop() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("this stack is empty, can't pop the element")
	}
	val := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return val, nil
}

func (stack *ArrayStack) Push(e interface{}) {
	stack.elements = append(stack.elements, e)
}

func (stack *ArrayStack) Clean() bool {
	stack.elements = make([]interface{}, 0, 10)
	return true
}
