package main

import (
	"fmt"
	"sort"
)

/*
	Input: text1 = "abcde", text2 = "ace"
	Output: 3
	Explanation: The longest common subsequence is "ace" and its length is 3.
 */
func main() {
	res := longestCommonSubsequence("abc", "def")
	for _, b := range res {
		fmt.Println(string(b))
	}
}

type step struct {
	i int
	j int
}

func longestCommonSubsequence(A string, B string) []byte {
	cache := make(map[step]bool)
	ids := make([]int, 0)
	dfs(A, B, 0, 0, cache, &ids)
	sort.Ints(ids)
	output := make([]byte, 0)
	for _, index := range ids {
		output = append(output, A[index])
	}
	return output
}

func dfs(A, B string, i, j int, cache map[step]bool, ids *[]int) {
	if i == len(A) || j == len(B) {
		return
	}

	if _, ok := cache[step{i, j}]; ok {
		return
	}

	if A[i] == B[j] {
		*ids = append(*ids, i)
		dfs(A, B, i+1, j+1, cache, ids)
	}else{
		dfs(A, B, i+1, j, cache, ids)
		dfs(A, B, i, j+1, cache, ids)
	}

	cache[step{i,j}] = true
}
