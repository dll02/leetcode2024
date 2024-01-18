package _50_evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if token == "*" {
				stack = append(stack, a*b)
			} else if token == "/" {
				stack = append(stack, a/b)
			} else if token == "+" {
				stack = append(stack, a+b)
			} else if token == "-" {
				stack = append(stack, a-b)
			}
		}

	}
	return stack[len(stack)-1]
}

func evalRPN2(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if token == "*" {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a*b)
		} else if token == "/" {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a/b)
		} else if token == "+" {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a+b)
		} else if token == "-" {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a-b)
		} else {
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}
	return stack[len(stack)-1]
}
