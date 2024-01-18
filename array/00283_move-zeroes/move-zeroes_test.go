package _0283_move_zeroes

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
	res []int
}

func Test_Func(t *testing.T) {
	questions := []question{
		{param{[]int{1, 2, 3, 0, 0, 0}},
			ans{[]int{1, 2, 3, 0, 0, 0}},
		},
		{
			param{[]int{1, 0, 2, 3}},
			ans{[]int{1, 2, 3, 0}},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		moveZeroes(qs.param.one)
		fmt.Printf("after merge sort res is %v \n", qs.param.one)
	}
}
