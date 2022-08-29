package main

import (
	"fmt"
)

/*
	Input: s = "leetcode", wordDict = ["leet","code"]
	Output: true
	Explanation: Return true because "leetcode" can be segmented as "leet code".
*/

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func wordBreak(s string, words []string) bool {
	cache := make([]int, len(s)+1)
	return dfs(s, words, cache)
}

func dfs(s string, words []string, cache []int) bool {
	if s == "" {
		return true
	}

	if cache[len(s)] != 0 {
		if cache[len(s)] == 1 {
			return true
		}else{
			return false
		}
	}

	flag := false
	for _, word := range words {
		if !isPrefix(s, word) {
			continue
		}
		flag = flag || dfs(s[len(word):], words, cache)
	}

	if flag{
		cache[len(s)] = 1
	}else{
		cache[len(s)] = 2
	}

	return flag
}

func isPrefix(s string, word string) bool {
	if len(word) > len(s) {
		return false
	}

	for i := 0; i < len(word); i++ {
		if s[i] != word[i] {
			return false
		}
	}

	return true
}