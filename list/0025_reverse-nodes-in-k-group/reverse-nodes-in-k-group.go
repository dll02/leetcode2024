package _025_reverse_nodes_in_k_group

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Println() {
	fmt.Print("[ ")
	for head != nil {
		fmt.Printf("%v,", head.Val)
		head = head.Next
	}
	fmt.Print("]\n")
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre, last := &ListNode{-1, head}, head
	ans := pre
	i := 1
	l, r := head, head

	for head != nil {
		r = head
		if r != nil {
			last = r.Next
		}

		if k == i {
			reversList(l, r)
			pre.Next = r
			l.Next = last
			pre = l
			head = l
			l = last

			i = 0
		}
		i++

		head = head.Next
	}
	return ans.Next
}

func reversList(start *ListNode, end *ListNode) (l *ListNode, r *ListNode) {
	if start == end || start == nil || end == nil {
		return
	}
	r = start
	l = end
	pre := start
	cur := start.Next
	for pre != end && cur != nil {
		last := cur.Next
		cur.Next = pre
		pre = cur
		cur = last
	}
	return l, r
}
