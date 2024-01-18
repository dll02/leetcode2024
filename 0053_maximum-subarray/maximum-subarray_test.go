package _053_maximum_subarray

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
	n int
}

func Test_Merge(t *testing.T) {
	questions := []question{
		//{param{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
		//	ans{6},
		//},
		{param{[]int{1}},
			ans{1},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		arr := maxSubArray(qs.param.one)
		fmt.Printf("ans is %v \n", arr)
	}
}
