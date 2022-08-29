package _59__Longest_Substring_with_At_Most_Two_Distinct_Characters_

/*
	Input: s = "eceba"
	Output: 3
	Explanation: The substring is "ece" which its length is 3.
*/
func lengthOfLongestSubstringTwoDistinct(s string) int {
	i, j, max := 0, 0, 0
	n := len(s)
	counter := make(map[byte]int)

	for j < n {
		if _, ok := counter[s[j]]; ok {
			counter[s[j]]++
			if j-i+1 > max {
				max = j - i + 1
			}
			j++
		} else {
			if len(counter) < 2 {
				counter[s[j]]++
				if j-i+1 > max {
					max = j - i + 1
				}
				j++
			}else{
				counter[s[i]]--
				if counter[s[i]] == 0 {
					delete(counter, s[i])
				}
				i++
			}
		}
	}
	return max
}
