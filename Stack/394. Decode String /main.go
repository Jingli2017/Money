package main

import (
	"fmt"
	"strings"
)

/*
	Input: s = "3[a]2[bc]"
	Output: "aaabcbc"
 */
func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}

func decodeString(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ']'{
			// pop up until [
			chars := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				popedChar := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				chars = append(chars, popedChar)
			}
			// pop up [
			stack = stack[:len(stack)-1]
			// pop up integer
			base := 1
			num := 0
			for len(stack) != 0 && isNumber(stack[len(stack)-1]){
				num = base * int(stack[len(stack)-1] - '0') + num
				base *= 10
				stack = stack[:len(stack)-1]
			}
			// push back
			for num != 0 {
				for j := len(chars)-1; j >= 0; j-- {
					stack = append(stack, chars[j])
				}
				num--
			}
		}else {
			stack = append(stack, s[i])
		}
	}

	var b strings.Builder
	for i := 0; i < len(stack); i++ {
		b.WriteByte(stack[i])
	}
	return b.String()
}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}
