package array_stack

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
