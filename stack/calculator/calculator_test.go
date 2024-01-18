package calculator

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
	res int
}

func TestCalculate(t *testing.T) {
	questions := []question{
		//{param{"12+211+36"},
		//	ans{6},
		//},
		{
			param{"(5+3)*10--9"},
			ans{71},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		res := calculate(qs.param.one)
		fmt.Printf("after merge sort res is %v \n", res)
	}

}
