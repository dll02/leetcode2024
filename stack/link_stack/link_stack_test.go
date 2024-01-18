package link_stack

import (
	"fmt"
	"testing"
)

func TestLinkStack(t *testing.T) {
	var s Stack
	s = New()
	s.Push(9)

	s.Push(100)
	s.Push(101)
	s.Push(1)
	fmt.Printf("%v\n", s)
	v, err := s.Pop()
	fmt.Printf("%v,%v \n", v, err)
	fmt.Printf("%v\n", s)
	s.Clean()
	fmt.Printf("%v\n", s)
	if s.Empty() {
		fmt.Println("is empty")
	}
}
