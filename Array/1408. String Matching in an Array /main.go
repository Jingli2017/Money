package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(stringMatching([]string{"mass","as","hero","superhero"}))
}

/*
	Input: words = ["mass","as","hero","superhero"]
	Output: ["as","hero"]
	Explanation: "as" is substring of "mass" and "hero" is substring of "superhero".
	["hero","as"] is also a valid answer.
 */
func stringMatching(words []string) []string {
	// sort by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	// check sub str
	set := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strStr(words[j], words[i]) != -1 {
				set[words[i]] = true
			}
		}
	}

	output := make([]string, 0)
	for key := range set {
		output = append(output, key)
	}

	return output
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

