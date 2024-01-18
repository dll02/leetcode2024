package _015_3sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := [][]int{}
	for i, x := range nums[:n-2] {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 {
			continue
		}
		if x+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s < 0 {
				j++
			} else if s > 0 {
				k--
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}
	return ans
}
