package with_array

type MyCircularDeque struct {
	front, rear int
	cap         int
	elements    []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{cap: k + 1, elements: make([]int, k+1)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = (this.front - 1 + this.cap) % this.cap
	this.elements[this.front] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.elements[this.rear] = value
	this.rear = (this.rear + 1) % this.cap
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.cap
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear - 1 + this.cap) % this.cap
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elements[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elements[(this.rear-1+this.cap)%this.cap]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == this.rear
}

func (this *MyCircularDeque) IsFull() bool {
	return this.front == (this.rear+1)%this.cap
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
