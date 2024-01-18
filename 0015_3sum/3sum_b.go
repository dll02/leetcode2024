package _015_3sum

import (
	"sort"
	"strconv"
)

func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	n := len(nums)
	// n-1 的范围数组 每次往里面放当前计算的值 替换 即将计算的值
	copy_nums := append([]int(nil), nums[:n-1]...)

	diffs := map[int][]int{}
	//var cur int
	for i := n - 2; i >= 0; i-- {
		cur := nums[i+1]
		// 不存在
		if _, ok := diffs[cur]; !ok {
			res := two_sum(copy_nums, -cur)
			if res != nil {
				diffs[cur] = res
			}
		}
		copy_nums[i] = cur

	}
	ans := [][]int{}
	// 怎么去重
	exists := map[string]bool{}
	for k, v := range diffs {
		arr := []int{k, v[0], v[1]}
		sort.Ints(arr)
		key := key(arr)
		if _, ok := exists[key]; !ok {
			ans = append(ans, arr)
			exists[key] = true
		}
	}
	return ans
}
func key(nums []int) string {

	key := ""
	for _, v := range nums {
		key = key + strconv.Itoa(v) + "_"
	}
	return key
}
func two_sum(nums []int, target int) []int {
	diffs := map[int]bool{}
	for _, v := range nums {
		if _, ok := diffs[target-v]; ok {
			return []int{v, target - v}
		}
		diffs[v] = true
	}
	return nil
}
