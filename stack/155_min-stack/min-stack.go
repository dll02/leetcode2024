package _55_min_stack

import "math"

/**
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。
*/
type MinStack struct {
	elements []int
	min      []int
}

func Constructor() MinStack {
	return MinStack{elements: []int{}, min: []int{math.MaxInt64}}
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)
	top := this.min[len(this.min)-1]
	this.min = append(this.min, min(val, top))

}

func (this *MinStack) Pop() {
	if len(this.elements) == 0 {
		return
	}
	this.elements = this.elements[:len(this.elements)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
