package _001_two_sum

func twoSum(nums []int, target int) []int {
	diffs := map[int]int{}

	for i, v := range nums {
		if p, ok := diffs[target-v]; ok {
			return []int{p, i}
		}
		diffs[v] = i
	}
	return nil
}
