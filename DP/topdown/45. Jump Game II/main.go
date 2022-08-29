package main

import (
	"fmt"
	"math"
)

/*
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
 */
func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
}

func jump(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	return dfs(nums, cache, 0)
}

func dfs(nums []int, cache []int, index int) int {
	n := len(nums)
	if index == n - 1 {
		return 0
	}

	if cache[index] != 0 {
		return cache[index]
	}

	furthest := getMin(index + nums[index], n-1)

	min := math.MaxInt32
	for i := index + 1; i <= furthest; i++ {
		step := dfs(nums, cache, i)
		if step < min{
			min = step
		}
	}

	if min != math.MaxInt32 {
		min++
	}

	cache[index] = min
	return min
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
