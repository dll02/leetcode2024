package search_neighbor_num

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type ListNode struct {
	Val   float64
	Index int
	Pre   *ListNode
	Next  *ListNode
}

func (node *ListNode) Println() {
	fmt.Print("[ ")
	for node != nil {
		fmt.Printf("%v[Index:%v],", node.Val, node.Index)
		node = node.Next
	}
	fmt.Print("]\n")
}

func (node *ListNode) DeleteNode() {
	if node.Next != nil {
		node.Next.Pre = node.Pre
	}
	if node.Pre != nil {
		node.Pre.Next = node.Next
	}
	node.Next = nil
	node.Pre = nil
}

type Result struct {
	Index int
	Val   float64
}

func main() {
	var n string
	fmt.Scanf("%s", &n)
	intN, _ := strconv.Atoi(n)

	//fmt.Printf("arg is %v \n", intN)
	lists := make([]*ListNode, intN)
	// read nums to listNode array
	for i := range lists {
		var str string
		fmt.Scanf("%s", &str)
		num, _ := strconv.Atoi(str)
		lists[i] = &ListNode{float64(num), i, nil, nil}
	}
	// sort  the listNode array
	sortLists := append([]*ListNode{}, lists...)
	sort.Slice(sortLists, func(i, j int) bool {
		return sortLists[i].Val < sortLists[j].Val
	})

	// build to Double Linked Lists
	head := &ListNode{-1, -1, nil, nil}
	tail := &ListNode{-1, -1, nil, nil}
	head.Next = tail
	tail.Pre = head
	cur := head
	for i := range sortLists {
		node := sortLists[i]
		node.Pre = cur
		cur.Next = node
		cur = node
	}
	cur.Next = tail
	tail.Pre = cur
	// from the max index node to get the |ai aj| answer
	//  then delete it from double linked list
	ans := make([]*Result, intN)
	for i := intN - 1; i > 0; i-- {
		node := lists[i]
		if (node.Pre.Index != -1 && math.Abs(node.Pre.Val-node.Val) <= math.Abs(node.Next.Val-node.Val)) || node.Next == tail {
			ans[node.Index] = &Result{node.Pre.Index, math.Abs(node.Pre.Val - node.Val)}
		} else {
			ans[node.Index] = &Result{node.Next.Index, math.Abs(node.Next.Val - node.Val)}
		}
		node.DeleteNode()
	}

	// print the answer
	i := 1
	for ; i < intN; i++ {
		fmt.Printf("%.0f %d\n", ans[i].Val, ans[i].Index+1)
	}
}

func searchNeighbor(arr []int, intN int) {
	for i := 1; i < intN; i++ {
		min := 0

		for j := 1; j < i; j++ {
			if math.Abs(float64(arr[i]-arr[j])) < math.Abs(float64(arr[i]-arr[min])) {
				min = j
			} else if math.Abs(float64(arr[i]-arr[j])) == math.Abs(float64(arr[i]-arr[min])) && arr[j] < arr[min] {
				min = j
			}
		}
		fmt.Printf("%.0f %d\n", math.Abs(float64(arr[i]-arr[min])), min+1)
	}
}
