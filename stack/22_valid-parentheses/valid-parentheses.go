package _2_valid_parentheses

func isValid(s string) bool {

	brackets := map[rune]rune{'}': '{', ')': '(', ']': '['}
	stack := []rune{}
	for _, ch := range s {
		if ch == '}' || ch == ')' || ch == ']' {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop != brackets[ch] {
				return false
			}
		} else {
			stack = append(stack, ch)
		}

	}

	return len(stack) == 0
}
