package main

func longestCommonSubsequence(A string, B string) int {
	m := len(A)
	n := len(B)

	memo := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		memo[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}

	return LCS(A, B, memo)
}

func LCS(A string, B string, memo [][]int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}

	if memo[len(A)][len(B)] != -1 {
		return memo[len(A)][len(B)]
	}

	lens := 0
	if A[0] == B[0] {
		lens = 1 + LCS(A[1:], B[1:], memo)
	} else {
		lens = getMax(LCS(A[1:], B, memo), LCS(A, B[1:], memo))
	}
	memo[len(A)][len(B)] = lens
	return lens
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
