package _2_valid_parentheses

import (
	"fmt"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	one string
}

type ans struct {
	res bool
}

func Test_Func(t *testing.T) {
	questions := []question{
		{param{"()"},
			ans{true},
		},
		{
			param{"()[]{}"},
			ans{true},
		},
		{
			param{"(]"},
			ans{false},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		res := isValid(qs.param.one)
		fmt.Printf("after merge sort res is %v \n", res)
	}
}
