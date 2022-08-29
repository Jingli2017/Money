package ___Longest_Substring_Without_Repeating_Characters_

/*
	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.
 */
func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool)
	i, j, max := 0, 0, 0
	n := len(s)

	for j < n {
		if _, ok := set[s[j]]; !ok {
			set[s[j]] = true
			if j - i + 1 > max {
				max = j - i + 1
			}
			j++
		}else{
			delete(set, s[i])
			i++
		}
	}

	return max
}