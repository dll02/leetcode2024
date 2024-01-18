package _109_corporate_flight_bookings

import (
	"fmt"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	one [][]int
	n   int
}

type ans struct {
	res []int
}

func Test_Merge(t *testing.T) {
	questions := []question{
		{param{[][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5},
			ans{[]int{10, 55, 45, 25, 25}},
		},
	}

	fmt.Println("===================leetcode 88 0088_merge-sorted-array===========================")
	for _, qs := range questions {
		fmt.Printf(" question is %v, answer is %v \n", qs.param, qs.ans)
		arr := corpFlightBookings(qs.param.one, qs.param.n)
		fmt.Printf("ans is %v \n", arr)
	}
}
