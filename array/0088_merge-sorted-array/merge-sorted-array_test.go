package leetcode

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
	m   int
	two []int
	n   int
}

type ans struct {
	res []int
}

func Test_Func(t *testing.T) {
	questions := []question{
		{param{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			ans{[]int{1, 2, 2, 3, 5, 6}},
		},
		{
			param{[]int{1, 0}, 1, []int{1}, 1},
			ans{[]int{1, 1}},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		merge(qs.param.one, qs.param.m, qs.param.two, qs.param.n)
		fmt.Printf("after merge sort res is %v \n", qs.param.one)
	}
}
