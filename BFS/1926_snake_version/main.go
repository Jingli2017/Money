package main

import "fmt"

/*
board1 =
[	['+', '+', '+', '+', '+', '+', '+', '0', '0'],
	['+', '+', '0', '0', '0', '0', '0', '+', '+'],
	['0', '0', '0', '0', '0', '+', '+', '0', '+'],
	['+', '+', '0', '+', '+', '+', '+', '0', '0'],
	['+', '+', '0', '0', '0', '0', '0', '0', '+'],
	['+', '+', '0', '+', '+', '0', '+', '0', '+']
]
start1_1 = (2, 0) # Expected output = (5, 2)
start1_2 = (0, 7) # Expected output = (0, 8)
start1_3 = (5, 2) # Expected output = (2, 0) or (5, 5)
start1_4 = (5, 5) # Expected output = (5, 7)
 */
func main() {
	board := [][]byte{
		{'+', '+', '+', '+', '+', '+', '+', '0', '0'},
		{'+', '+', '0', '0', '0', '0', '0', '+', '+'},
		{'0', '0', '0', '0', '0', '+', '+', '0', '+'},
		{'+', '+', '0', '+', '+', '+', '+', '0', '0'},
		{'+', '+', '0', '0', '0', '0', '0', '0', '+'},
		{'+', '+', '0', '+', '+', '0', '+', '0', '+'},
	}

	entrance := [][]int{{2,0}, {0,7}, {5,2}, {5,5}}

	for _, en := range entrance{
		duplicate := make([][]byte, len(board))
		for i := 0; i < len(board); i++ {
			duplicate[i] = make([]byte, len(board[0]))
			copy(duplicate[i], board[i])
		}
		fmt.Println(nearestExit(duplicate, en))
	}
}

var dirs = [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}}

type point struct {
	x int
	y int
}

func nearestExit(board [][]byte, entrance []int) []int {
	m := len(board)
	n := len(board[0])
	q := make([]point, 0)
	q = append(q, point{entrance[0], entrance[1]})
	board[entrance[0]][entrance[1]] = '+'

	for len(q) != 0 {
		size := len(q)
		for k := 0; k < size; k++ {
			curr := q[0]
			q = q[1:]

			for i := 0; i < 4; i++ {
				newx := curr.x + dirs[i][0]
				newy := curr.y + dirs[i][1]
				// out of bound
				if newx < 0 || newx == m || newy < 0 || newy == n { continue }
				// visited before or it is wall
				if board[newx][newy] == '+' { continue }
				// on the boarder
				if newx == 0 || newx == m-1 || newy == 0 || newy == n-1 {
					return []int{newx, newy}
				}
				// mark visited and put into queue
				board[newx][newy] = '+'
				q = append(q, point{newx, newy})
			}
		}
	}
	return nil
}

