package main

import "fmt"

/*
	Input: haystack = "hello", needle = "ll"
	Output: 2
 */
func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if i + len(needle) > len(haystack) {
			return -1
		}
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle) - 1 {
				return i
			}
		}
	}

	return -1
}
