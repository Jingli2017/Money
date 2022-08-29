package _00__Number_of_Islands_

/*
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
 */
var dirs = [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}}
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1'{
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i, j int){
	m := len(grid)
	n := len(grid[0])

	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}

	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	for _, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]
		dfs(grid, x, y)
	}
}
