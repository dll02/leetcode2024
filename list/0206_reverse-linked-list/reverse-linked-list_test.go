package _206_reverse_linked_list

import (
	"fmt"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	list *ListNode
}

type ans struct {
	list *ListNode
}

func Test_Func(t *testing.T) {
	questions := []question{
		{
			param{&ListNode{10, &ListNode{8, &ListNode{9, &ListNode{7, nil}}}}},
			ans{&ListNode{7, &ListNode{9, &ListNode{8, &ListNode{10, nil}}}}},
		},
		{
			param{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
			ans{&ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Println(" old listNode ")
		qs.param.list.Println()
		head := reverseList(qs.param.list)
		fmt.Println(" after reverseList list is")
		head.Println()
	}
}
