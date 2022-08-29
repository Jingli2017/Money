package main

import "fmt"

func main() {
	fmt.Println(findLength([]int{1,2,3,2,1}, []int{3,2,1,4,7}))
}

func findLength(A []int, B []int) int {
	m, n := len(A), len(B)
	max := 0
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	dfs(A, B, 0, 0, cache, &max)
	return max
}

func dfs(A, B []int, i, j int, cache [][]int, max *int) int {
	if i == len(A) || j == len(B) { return 0 }
	if cache[i][j] != -1 {
		return cache[i][j]
	}

	length := 0
	if A[i] == B[j] {
		length = 1 + dfs(A, B, i+1, j+1, cache, max)
		if length > *max {
			*max = length
		}
	}

	dfs(A, B, i, j+1, cache, max)
	dfs(A, B, i+1, j, cache, max)

	cache[i][j] = length
	return length
}
