package _1_merge_two_sorted_lists

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur := &ListNode{-1, nil}
	head := cur
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return head.Next
}
