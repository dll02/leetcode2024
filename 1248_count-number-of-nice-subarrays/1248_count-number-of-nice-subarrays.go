package _248_count_number_of_nice_subarrays

func numberOfSubarrays(nums []int, k int) int {

	n := len(nums)
	s := make([]int, n+1)
	counts := make([]int, n+1)
	counts[0] = 1
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + nums[i]%2
		counts[s[i+1]]++
	}

	ans := 0
	for i := 0; i < n; i++ {
		if s[i+1]-k >= 0 {
			ans += counts[s[i+1]-k]
		}
	}

	return ans
}
