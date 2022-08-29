package _29__Longest_Increasing_Path_in_a_Matrix

/*
	matrix = [[9,9,4],[6,6,8],[2,1,1]]
	9 9 4
	6 6 8
	2 1 1
 */

var dirs = [][]int{{1,0}, {-1,0}, {0,-1}, {0,1}}
func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}

	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			length := dfs(matrix, i, j, cache)
			if length > max {
				max = length
			}
		}
	}

	return max
}

func dfs(matrix [][]int, i, j int, cache [][]int) int {
	if cache[i][j] != 0 { return cache[i][j] }

	m := len(matrix)
	n := len(matrix[0])

	max := 0
	for _, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]

		if x < 0 || x == m || y < 0 || y == n {
			continue
		}

		if matrix[i][j] >= matrix[x][y] {
			continue
		}

		length := dfs(matrix, x, y, cache)
		if length > max {
			max = length
		}
	}

	cache[i][j] = max + 1
	return max + 1
}
