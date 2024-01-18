package _025_reverse_nodes_in_k_group

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
	k    int
}

type ans struct {
	list *ListNode
}

func Test_Func(t *testing.T) {
	questions := []question{
		//{
		//	param{&ListNode{10, &ListNode{8, &ListNode{9, &ListNode{7, nil}}}}, 2},
		//	ans{&ListNode{8, &ListNode{10, &ListNode{7, &ListNode{9, nil}}}}},
		//},
		{
			param{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}, 3},
			ans{&ListNode{3, &ListNode{2, &ListNode{1, &ListNode{6, &ListNode{5, &ListNode{4, &ListNode{7, nil}}}}}}}},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Println(" old listNode ")
		qs.param.list.Println()
		head := reverseKGroup(qs.param.list, qs.param.k)
		fmt.Println(" after reverseList list is")
		head.Println()
	}
}
