package main

import "fmt"

func main() {
	fmt.Println(minKnightMoves(2, 1))
}

var dirs = [8][2]int{
	{-2, 1},
	{-1, 2},
	{1, 2},
	{2, 1},
	{2, -1},
	{1, -2},
	{-1, -2},
	{-2, -1},
}

type position struct {
	x int
	y int
}

func minKnightMoves(x int, y int) int {
	visited := make(map[position]bool)
	visited[position{x:0, y:0}] = true
	q := make([]position, 0)
	q = append(q, position{x:0, y:0})
	steps := 0

	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			curr := q[0]
			q = q[1:]

			if curr.x == x && curr.y == y {return steps}

			for _, dir := range dirs {
				nextX := curr.x + dir[0]
				nextY := curr.y + dir[1]
				newPosition := position{x:nextX, y:nextY}
				if _, ok := visited[newPosition]; ok {
					continue
				}
				visited[newPosition] = true
				q = append(q, newPosition)
			}
		}
		steps++
	}

	return steps
}
