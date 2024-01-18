package _011_container_with_most_water

func maxArea(height []int) int {
	ans := 0
	n := len(height)
	i, k := 0, n-1
	for i < k {
		ans = max(ans, min(height[i], height[k])*(k-i))
		if height[i] == height[k] {
			i++
			k--
		} else if height[i] < height[k] {
			i++
		} else {
			k--
		}
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
