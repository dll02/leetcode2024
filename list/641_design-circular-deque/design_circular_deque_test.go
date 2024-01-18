package _41_design_circular_deque

import "testing"

func TestConstructor(t *testing.T) {
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
	k := 10
	obj := Constructor(k)
	obj.InsertFront(12)
	obj.InsertLast(33)

}
