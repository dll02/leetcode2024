package _066_plus_one

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

func plusOne1(digits []int) []int {
	// 模拟计算器 carry 进位
	carry := 1
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if carry == 1 && digits[i] == 9 {
			digits[i] = 0
			carry = 1
		} else {
			// <9
			if carry == 1 {
				digits[i] += 1
			}
			carry = 0
		}
	}
	ans := []int{1}
	if carry == 1 {
		ans = append(ans, digits...)
	} else {
		ans = digits
	}
	return ans
}

//  超出 int bit 位数开始报错
func plusOne2(digits []int) []int {
	num := 0
	for _, v := range digits {
		num = num*10 + v
	}
	num = num + 1
	ans := []int{}
	for num > 0 {
		ans = append(ans, num%10)
		num = num / 10
	}

	i, j := 0, len(ans)-1
	for i != j && i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}

	return ans
}
