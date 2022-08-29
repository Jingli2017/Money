package main

import "fmt"

/*
	Input: s = "ADOBECODEBANC", t = "ABC"
	Output: "BANC"
	Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
 */
func main() {
	fmt.Println(minWindow("aa", "aa"))
}

func minWindow(A string, B string) string {
	if A == "" || B == "" || len(B) > len(A) { return "" }
	BCount := make(map[byte]int)
	for i := 0; i < len(B); i++ {
		BCount[B[i]]++
	}

	res := make([]int, 2)
	left, right, meet, min := 0, 0, 0, -1
	window := make(map[byte]int)

	for right < len(A) {
		window[A[right]]++

		if window[A[right]] == BCount[A[right]]{
			meet++
		}

		for meet == len(BCount) && left <= right {
			if min == -1 || right - left + 1 < min {
				min = right - left + 1
				res[0], res[1] = left, right
			}
			window[A[left]]--
			if count, ok := BCount[A[left]]; ok && window[A[left]] < count {
				meet--
			}
			left++
		}
		right++
	}

	if min == -1 {
		return ""
	}
	return A[res[0]:res[1]+1]
}
