package _015_3sum

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
	ans [][]int
}

func Test_Merge(t *testing.T) {
	questions := []question{
		//{param{[]int{0, 0, 0}},
		//	// [[-1,-1,2],[-1,0,1]]
		//	ans{[][]int{{0, 0, 0}}},
		//},
		{param{[]int{-1, 0, 1, 2, -1, -4}},
			// [[-1,-1,2],[-1,0,1]]
			ans{[][]int{{-1, -1, 2}, {-1, 0, 1}}},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		arr :=
			threeSum(qs.param.one)
		fmt.Printf("ans is %v \n", arr)
	}
}
