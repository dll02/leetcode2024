package with_slice

// 底层用切片实现
type MyCircularDeque struct {
	cap      int
	elements []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{cap: k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.elements = append([]int{value}, this.elements...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.elements = append(this.elements, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.elements = this.elements[1:]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.elements = this.elements[:len(this.elements)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elements[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elements[len(this.elements)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.elements) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return len(this.elements) == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
