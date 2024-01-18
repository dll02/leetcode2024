package _026_remove_duplicates_from_sorted_array

import (
	"fmt"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	one []int
}

type ans struct {
	n   int
	res []int
}

func Test_Merge(t *testing.T) {
	questions := []question{
		{param{[]int{1, 1, 2}},
			ans{2, []int{1, 2}},
		},
		{
			param{[]int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 4}},
			ans{5, []int{0, 1, 2, 3, 4}},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		n := removeDuplicates(qs.param.one)
		fmt.Printf("after n is %v res is %v \n", n, qs.param.one)
	}
}
