package _1_merge_two_sorted_lists

import (
	"fmt"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	list1 *ListNode
	list2 *ListNode
}

type ans struct {
	list *ListNode
}

func Test_Func(t *testing.T) {
	questions := []question{
		{
			param{&ListNode{1, &ListNode{2, &ListNode{4, &ListNode{7, nil}}}},
				&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{9, nil}}}}},
			ans{},
		},
		{
			param{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}},
				&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{6, &ListNode{8, &ListNode{9, &ListNode{17, nil}}}}}}},
			},
			ans{},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Println(" old listNode ")
		qs.param.list1.Println()
		qs.param.list2.Println()
		head := mergeTwoLists(qs.param.list1, qs.param.list2)
		fmt.Println(" after reverseList list is")
		head.Println()
	}
}
