package main

/*
Input: grid = [
[0,0,1,0,0,0,0,1,0,0,0,0,0],
[0,0,0,0,0,0,0,1,1,1,0,0,0],
[0,1,1,0,1,0,0,0,0,0,0,0,0],
[0,1,0,0,1,1,0,0,1,0,1,0,0],
[0,1,0,0,1,1,0,0,1,1,1,0,0],
[0,0,0,0,0,0,0,0,0,0,1,0,0],
[0,0,0,0,0,0,0,1,1,1,0,0,0],
[0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.
*/
func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			temp := dfs(grid, i, j)
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0{
		return 0
	}

	grid[i][j] = 0 // mark explored

	top := dfs(grid, i-1, j)
	down := dfs(grid, i+1, j)
	left := dfs(grid, i, j-1)
	right := dfs(grid, i, j+1)

	return 1 + top + down + left + right
}
