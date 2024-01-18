package _0283_move_zeroes

func moveZeroes(nums []int) {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}

}

func moveZeroes2(nums []int) {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if p != i {
				nums[p] = nums[i]
			}
			p++
		}
	}

	for ; p < len(nums); p++ {
		nums[p] = 0
	}
}
func moveZeroes3(nums []int) {
	p, q := 0, 0
	for q < len(nums) {
		if nums[q] != 0 {
			if q != p {
				nums[p] = nums[q]
			}
			p++
		}
		q++
	}

	for p < len(nums) {
		nums[p] = 0
		p++
	}
}
