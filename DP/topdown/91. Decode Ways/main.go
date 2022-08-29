package main

import (
	"fmt"
	"strconv"
)

/*
Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
*/
func main() {
	fmt.Println(numDecodings("226"))
}

func numDecodings(s string) int {
	cache := make(map[int]int)
	return dfs(s, 0, cache)
}

func dfs(s string, index int, cache map[int]int) int {
	if index == len(s) {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	if _, ok := cache[index]; ok {
		return cache[index]
	}

	ans := dfs(s, index + 1, cache)
	if index + 2 <= len(s) && isLessThan26(s[index: index+2]){
		ans += dfs(s, index + 2, cache)
	}

	cache[index] = ans
	return ans
}

func isLessThan26(s string) bool {
	num, _ := strconv.Atoi(s)
	return num <= 26
}

