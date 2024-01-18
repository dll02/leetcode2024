package _41_design_circular_deque

type node struct {
	val  int
	next *node
	pre  *node
}
type MyCircularDeque struct {
	len  int
	cap  int
	head *node
	tail *node
}

func Constructor(k int) MyCircularDeque {

	return MyCircularDeque{cap: k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &node{val: value}
	if this.IsEmpty() {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.pre = node
		this.head = node
	}
	this.len++
	return true

}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &node{val: value}
	if this.IsEmpty() {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		node.pre = this.tail
		this.tail = node
	}
	this.len++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = this.head.next
	if this.head != nil {
		this.head.pre = nil
	}
	this.len--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = this.tail.pre
	if this.tail != nil {
		this.tail.next = nil
	}
	this.len--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.cap
}
