package _55_min_stack

import (
	"fmt"
	"testing"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.push(10)
	obj.push(1)
	obj.push(4)
	min := obj.getMin()
	fmt.Println(min)
}
