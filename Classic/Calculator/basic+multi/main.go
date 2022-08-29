package main

import "fmt"

func main() {
	s := "2*(5+5*2)/3+(6/2+8)"
	fmt.Println(calculate(s))
}

// s = "2*(5+5*2)/3+(6/2+8)"
func calculate(s string) int {
	q := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		q = append(q, s[i])
	}
	q = append(q, '+')
	return calImpl(&q)
}

func calImpl(q *[]uint8) int {
	stack := make([]int, 0)
	var sign uint8 = '+'
	operand := 0

	for len(*q) != 0 {
		c := (*q)[0]
		*q = (*q)[1:]

		if c == ' ' { continue }
		if c >= '0' && c <= '9' {
			operand = operand*10 + int(c - '0')
		}else if c == '('{
			operand = calImpl(q)
		}else{ // +,-,*,/
			if sign == '+' {
				stack = append(stack, operand)
			}else if sign == '-' {
				stack = append(stack, -operand)
			}else if sign == '*' {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop * operand)
			}else { // divide
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop / operand)
			}

			if c == ')' { break }
			operand = 0
			sign = c
		}
	}

	ans := 0
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans += pop
	}

	return ans
}
