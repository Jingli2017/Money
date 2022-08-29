package main

import "fmt"

/*
	Input: nums = [10,9,2,5,3,7,101,18]
	Output: 4
	Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
*/
func main() {
	fmt.Println(lengthOfLIS([]int{1,2}))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	max := 0
	cache := make([]int, n)
	for i := 0; i < len(nums); i++ {
		temp := dfs(nums, i, cache)
		if temp > max {
			max = temp
		}
	}
	return max
}

func dfs(nums []int, i int, cache []int) int {
	n := len(nums)
	if i == n {
		return 0
	}

	if cache[i] != 0 {
		return cache[i]
	}

	max := 0
	for j := i + 1; j < n; j++ {
		if nums[i] < nums[j] {
			temp := dfs(nums, j, cache)
			if max < temp {
				max = temp
			}
		}
	}

	cache[i] = max + 1
	return max + 1
}
