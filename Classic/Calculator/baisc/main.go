package main

import "fmt"

func calculate(s string) int {
	q := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		q = append(q, s[i])
	}
	q = append(q, '+')
	return impl(&q)
}

func impl(q *[]uint8) int {
	var sign uint8 = '+'
	operand := 0
	nums := make([]int, 0)
	for len(*q) != 0 {
		c := (*q)[0]
		*q = (*q)[1:]
		if c == ' ' {
			continue
		}else if isDigit(c) {
			operand = 10 * operand + int(c - '0')
		}else if c == '(' {
			operand = impl(q)
		}else{
			if sign == '+' {
				nums = append(nums, operand)
			}else if sign == '-'{
				nums = append(nums, -1 * operand)
			}
			if c == ')' {
				break
			}
			sign = c
			operand = 0
		}
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func main() {
	fmt.Println(calculate("(1+2)"))
}