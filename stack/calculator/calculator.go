package calculator

import (
	"fmt"
	"strconv"
)

// 先计算成后缀表达式
// 后缀表达式
func calculate(s string) int {
	tokens := []string{}
	ops := []string{}
	num_start := false
	need_zero := false
	num := 0
	for _, ch := range s {
		// 处理 num
		if '0' <= ch && ch <= '9' {
			num = num*10 + int(ch-'0')
			num_start = true
			continue
		} else if num_start {
			// 处理到非数据部分
			tokens = append(tokens, strconv.Itoa(num))
			num_start = false
			need_zero = false
			num = 0
		}
		// 处理其他符号
		switch ch {
		case ' ': // ' '
			continue
		case '(': // (
			need_zero = true
			ops = append(ops, string(ch))
			continue
		case ')':
			for len(ops) > 0 && ops[len(ops)-1] != string('(') {
				tokens = append(tokens, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			// pop (
			ops = ops[:len(ops)-1]
			need_zero = false
			continue
		default:
			if need_zero {
				tokens = append(tokens, strconv.Itoa(0))
				need_zero = false
			}
			for len(ops) > 0 && rank(ops[len(ops)-1]) >= rank(string(ch)) {
				tokens = append(tokens, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, string(ch))
			need_zero = true
		}
	}
	// num_start 入栈
	if num_start {
		tokens = append(tokens, strconv.Itoa(num))
	}
	// ops 入栈
	if len(ops) > 0 {
		tokens = append(tokens, ops[len(ops)-1])
	}
	fmt.Println(tokens)
	// 计算后缀表达式
	return evalRPN(tokens)
}

func rank(ch string) int {
	if ch == string('*') || ch == string('/') {
		return 2
	} else if ch == string('+') || ch == string('-') {
		return 1
	} else {
		return 0
	}
}

func evalRPN(tokens []string) int {
	stack := []int{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			res := clac(num1, num2, token)
			stack = append(stack, res)

		}
	}
	return stack[len(stack)-1]
}

func clac(num1 int, num2 int, ch string) int {
	switch ch {
	case string('+'):
		return num1 + num2
	case string('-'):
		return num1 - num2
	case string('*'):
		return num1 * num2
	case string('/'):
		return num1 / num2
	default:
		return -1
	}
}
