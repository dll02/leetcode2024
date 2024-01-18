package link_stack

import "errors"

type Stack interface {
	//	push(e): Add e at the top of the (implicit) stack
	//pop(): Remove and return the top element of the stack
	//
	//empty(): Return the Boolean value true just in case the stack is empty.
	//top(): Return the top element of that stack without removing it.
	Empty() bool
	Top() (interface{}, error)
	Pop() (interface{}, error)
	Push(e interface{})
	Clean() bool
}

type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkStack struct {
	top *ListNode
	len int
}

func New() *LinkStack {
	return &LinkStack{nil, 0}
}

func (stack *LinkStack) Empty() bool {
	return stack.len == 0
}

func (stack *LinkStack) Top() (interface{}, error) {
	if stack.len == 0 {
		return nil, errors.New("this stack is empty, can't top the element")
	}
	return stack.top.val, nil
}

func (stack *LinkStack) Pop() (interface{}, error) {
	if stack.len == 0 {
		return nil, errors.New("this stack is empty, can't pop the element")
	}
	val := stack.top.val
	stack.top = stack.top.next
	return val, nil
}

func (stack *LinkStack) Push(e interface{}) {
	node := &ListNode{e, nil}
	if stack.top == nil {
		stack.top = node
	} else {
		node.next = stack.top
		stack.top = node
	}
	stack.len++
}

func (stack *LinkStack) Clean() bool {
	stack.top = nil
	stack.len = 0
	return true
}
