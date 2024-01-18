package search_neighbor_num

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
	n   int
}

type ans struct {
	res [][]int
}

func Test_Func(t *testing.T) {
	questions := []question{
		//{param{[]int{1, 5, 3}, 3},
		//	ans{[][]int{{4, 1}, {2, 1}}},
		//},
		//{param{[]int{2, 3, 4, 7, 1, 8}, 6},
		//	ans{[][]int{{1, 1}, {1, 2}, {3, 3}, {1, 1}, {1, 4}}},
		//},
		{
			param{[]int{4, 5, 6, 1, 2, 3, 7, 8, 9, 10}, 10},
			ans{[][]int{{1, 1},
				{1, 2},
				{3, 1},
				{1, 4},
				{1, 5},
				{1, 3},
				{1, 7},
				{1, 8},
				{1, 9}}},
		},
		{
			param{
				[]int{-1000000000, 0, -999999999, 999999999, 1000000000}, 5},
			ans{[][]int{}},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		searchNeighbor(qs.param.one, qs.param.n)
		fmt.Printf("after merge sort res is %v \n", qs.param.one)
	}
}
