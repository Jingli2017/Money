package main

import "fmt"

/*
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
*/

/*
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/
func main() {
	fmt.Println(canJump([]int{2, 1}))
}

const (
	UN = iota
	Y
	N
)

func canJump(nums []int) bool {
	cache := make([]int, len(nums))
	return dfs(nums, cache, 0)
}

func dfs(nums []int, cache []int, index int) bool {
	if index == len(nums)-1 {
		return true
	}
	if cache[index] != UN {
		if cache[index] == Y {
			return true
		}else {
			return false
		}
	}

	farestIndex := getMin(len(nums)-1, nums[index]+index)
	for i := index + 1; i <= farestIndex; i++ {
		if dfs(nums, cache, i) {
			cache[index] = Y
			return true
		}
	}

	cache[index] = N
	return false
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
