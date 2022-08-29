package main

import "fmt"

/*
	Input: arr = [4,2,3,0,3,1,2], start = 5
	Output: true
	Explanation:
	All possible ways to reach at index 3 with value 0 are:
	index 5 -> index 4 -> index 1 -> index 3
	index 5 -> index 6 -> index 4 -> index 1 -> index 3
*/

/*
	[3,0,2,1,2]
	[2]
	[0, 4]
	[3]
	X
	2
 */

/*
	[4,2,3,0,3,1,2]
	5
 */
func main() {
	fmt.Println(canReach([]int{0,1}, 1))
}

func canReach(arr []int, start int) bool {
	q := make([]int, 0)
	q = append(q, start)

	visited := make(map[int]bool)
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		if _, ok := visited[curr]; ok {
			continue
		}

		if arr[curr] == 0 {
			return true
		}

		visited[curr] = true

		if curr + arr[curr] < len(arr) {
			q = append(q, curr + arr[curr])
		}
		if curr - arr[curr] >= 0 {
			q = append(q, curr - arr[curr])
		}
	}

	return false
}