package _053_maximum_subarray

import "math"

// 找差分 最大间隔值 s[i]-s[j] 中 s[j] 最小
func maxSubArray(nums []int) int {
	n := len(nums)
	sums := make([]int, n+1)
	sums[0] = nums[0]
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	pre_min := make([]int, n+1)
	pre_min[0] = sums[0]
	for i := 1; i <= n; i++ {
		pre_min[i] = int(math.Min(float64(sums[i]), float64(pre_min[i-1])))
	}

	ans := -math.MaxInt
	for i := 1; i <= n; i++ {

		ans = int(math.Max(float64(ans), float64(sums[i]-pre_min[i-1])))
	}
	return ans
}
