package main

func calculate(s string) int {
	s += "+" // to push last operand
	stack := make([]int, 0)

	operand := 0
	var sign uint8 = '+'
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' { continue }
		if s[i] >= '0' && s[i] <= '9' {
			operand = operand * 10 + int(s[i] - '0')
		}else{ // s[i] == +, -, *, /
			if sign == '+' {
				stack = append(stack, operand)
			}else if sign == '-' {
				stack = append(stack, -operand)
			}else if sign == '*' {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop * operand)
			}else { // divide ops
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop / operand)
			}
		}
		sign = s[i]
		operand = 0
	}

	ans := 0
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans += pop
	}

	return ans
}
