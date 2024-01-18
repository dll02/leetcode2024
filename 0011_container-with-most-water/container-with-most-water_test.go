package _011_container_with_most_water

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
	ans int
}

func Test_Merge(t *testing.T) {
	questions := []question{
		{param{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			ans{49},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		arr := maxArea(qs.param.one)
		fmt.Printf("ans is %v \n", arr)
	}
}
