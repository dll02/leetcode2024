package _167_two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	diffs := map[int]int{}
	for i, v := range numbers {
		if p, ok := diffs[target-v]; ok {
			return []int{p + 1, i + 1}
		}
		diffs[v] = i
	}
	return nil
}
