package _026_remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[p] = nums[i]
			p++
		}
	}
	return p
}
