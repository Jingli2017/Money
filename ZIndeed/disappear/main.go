package main

import "fmt"

/*
While your players are waiting for a game, you've developed a solitaire game for the players to pass the time with.
The player is given an NxM board of tiles from 0 to 9 like this:
4 4 4 4
5 5 5 4
2 5 7 5
The player selects one of these tiles, and that tile will disappear,
along with any tiles with the same number that are connected with that tile
(up, down, left, or right), and any tiles with the same number connected with those, and so on. For example,
if the 4 in the upper left corner is selected, these five tiles disappear
*/
/*
	disappear(grid1, 0, 0) => 5
	disappear(grid1, 1, 1) => 4
	disappear(grid1, 1, 0) => 4
*/
func main() {
	grid := [][]int{
		{4, 4, 4, 4},
		{5, 5, 5, 4},
		{2, 5, 7, 5},
	}
	cases := [][]int{
		{0, 0},
		{1, 1},
		{1, 0},
	}

	for _, point := range cases {
		fmt.Println(disappear(grid, point[0], point[1]))
	}
}

func disappear(grid [][]int, x, y int) int {
	return dfs(grid, x, y, grid[x][y])
}

func dfs(grid [][]int, x, y int, num int) int {
	if x < 0 || x == len(grid) || y < 0 || y == len(grid[0]) {
		return 0
	}

	if grid[x][y] == -1 || grid[x][y] != num {
		return 0
	}

	grid[x][y] = -1

	top := dfs(grid, x-1, y, num)
	down := dfs(grid, x+1, y, num)
	left := dfs(grid, x, y-1, num)
	right := dfs(grid, x, y+1, num)

	return 1 + top + down + left + right
}
