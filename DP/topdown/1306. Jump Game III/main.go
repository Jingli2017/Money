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
	2
 */
/*
	[4,2,3,0,3,1,2]
	5
 */
func main() {
	fmt.Println(canReach([]int{4,2,3,0,3,1,2}, 5))
}

func canReach(arr []int, start int) bool {
	cache := make(map[int]bool)
	return dfs(arr, cache, start)
}

func dfs(arr []int, cache map[int]bool, step int) bool {
	if step < 0 || step >= len(arr) {
		return false
	}

	if _, ok := cache[step]; ok {
		return false
	}

	if arr[step] == 0 {
		return true
	}

	cache[step] = true

	return dfs(arr, cache, step+arr[step]) || dfs(arr, cache, step-arr[step])
}