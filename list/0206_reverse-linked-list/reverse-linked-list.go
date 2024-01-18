package _206_reverse_linked_list

import "fmt"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

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

func reverseList(head *ListNode) *ListNode {

	var pre *ListNode

	for head != nil {
		next := head.Next
		head.Next = pre

		pre = head
		head = next
	}
	return pre
}
