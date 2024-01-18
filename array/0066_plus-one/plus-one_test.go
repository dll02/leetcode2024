package _066_plus_one

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
		{param{[]int{9, 9, 9}},
			ans{[]int{1, 0, 0, 0}},
		},
		{
			param{[]int{1, 0, 2, 3}},
			ans{[]int{1, 0, 2, 4}},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		ans := plusOne(qs.param.one)
		fmt.Printf("after func res is %v \n", ans)
	}
}
