package leetcode

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Printf("%v", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n; i > 0; i-- {
		if m > 0 && n > 0 && nums1[m-1] >= nums2[n-1] {
			nums1[i-1] = nums1[m-1]
			m--
		} else if m > 0 && n > 0 {
			nums1[i-1] = nums2[n-1]
			n--
		}

	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n; i > 0; i-- {
		if n == 0 || (m > 0 && nums1[m-1] > nums2[n-1]) {
			nums1[i-1] = nums1[m-1]
			m--
		} else if n > 0 {
			nums1[i-1] = nums2[n-1]
			n--
		}
	}
}
func merge1(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	size := m + n
	ret := make([]int, size)
	l := 0
	r := 0
	for i := 0; i < size; i++ {
		if r >= n || (l < m && nums1[l] < nums2[r]) {
			ret[i] = nums1[l]
			fmt.Printf("%d:%d\n", l, nums1[l])
			l++
		} else {
			ret[i] = nums2[r]
			fmt.Printf("%d:%d\n", r, nums2[r])
			r++
		}
	}
	for i := 0; i < size; i++ {
		nums1[i] = ret[i]
	}

}
