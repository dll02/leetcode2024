package _015_3sum

import "sort"

func threeSum3(nums []int) [][]int {
	// 排序后双指针
	sort.Ints(nums)
	ans := [][]int{}
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoans := twosum(nums, i+1, -nums[i])
		if twoans != nil {
			for _, v := range twoans {
				ans = append(ans, []int{nums[i], v[0], v[1]})
			}
		}
	}
	return ans
}

func twosum(nums []int, start int, target int) [][]int {
	ans := [][]int{}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		j := len(nums) - 1
		for ; i < j && nums[i]+nums[j] > target; j-- {

		}
		if i == j {
			break
		}
		if nums[i]+nums[j] == target {
			ans = append(ans, []int{nums[i], nums[j]})
		}
	}
	return ans
}
