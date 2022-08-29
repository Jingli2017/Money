package _21__Maximal_Square_

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
				}
			}else{
				if matrix[i][j] == '1' {
					dp[i][j] = getMin(dp[i-1][j-1], getMin(dp[i-1][j], dp[i][j-1])) + 1
				}
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}
		}
	}

	return max * max
}

func getMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}
