package main

import "fmt"

/*
[["+",".","+","+","+","+","+"],
 ["+",".","+",".",".",".","+"],
 ["+",".","+",".","+",".","+"],
 ["+",".",".",".","+",".","+"],
 ["+","+","+","+","+",".","+"]]
[0,1]
 */
func main() {
	board := [][]byte{
		{'+','.','+','+','+','+','+'},
		{'+','.','+','.','.','.','+'},
		{'+','.','+','.','+','.','+'},
		{'+','.','.','.','+','.','+'},
		{'+','+','+','+','+','.','+'},
	}

	fmt.Println(nearestExit(board, []int{0,1}))
}

var dirs = [][]int{{1,0}, {-1, 0}, {0,1}, {0,-1}}

type point struct {
	x int
	y int
}

func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	step := 0
	q := make([]point, 0)
	q = append(q, point{entrance[0], entrance[1]})
	visited[entrance[0]][entrance[1]] = true

	for len(q) != 0 {
		size := len(q)
		step++
		for k := 0; k < size; k++ {
			current := q[0]
			q = q[1:]

			// explore four directions
			for i := 0; i < 4; i++ {
				newx := dirs[i][0] + current.x
				newy := dirs[i][1] + current.y
				if isOutOfBound(maze, newx, newy) { continue }
				if maze[newx][newy] == '+' || visited[newx][newy] { continue }
				if isBoarder(maze, newx, newy) {
					return step
				}

				visited[newx][newy] = true
				q = append(q, point{newx, newy})
			}
		}
	}
	return -1
}

func isOutOfBound(maze [][]byte, x int, y int) bool {
	m := len(maze)
	n := len(maze[0])

	if x < 0 || x == m || y < 0 || y == n {
		return true
	}

	return false
}

func isBoarder(maze [][]byte, x int, y int) bool {
	m := len(maze)
	n := len(maze[0])

	if x == 0 || x == m-1 || y == 0 || y == n-1 {
		return true
	}

	return false
}

