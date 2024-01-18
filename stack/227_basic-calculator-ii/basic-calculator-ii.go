package _27_basic_calculator_ii

// just only + - * /
func calculate(s string) int {
	// num prepare
	// push + - => sign*num into stack
	//  * / => cue result int last stack num
	num := 0
	n := len(s)
	stack := []int{}
	ans := 0
	presign := '+'
	for i, ch := range s {
		digint := '0' <= ch && ch <= '9'
		if digint {
			num = num*10 + int(ch-'0')
		}
		if !digint && ch != ' ' || i == n-1 {
			switch presign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -1*num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			num = 0
			presign = ch
		}
	}
	for _, v := range stack {
		ans += v
	}

	return ans
}
