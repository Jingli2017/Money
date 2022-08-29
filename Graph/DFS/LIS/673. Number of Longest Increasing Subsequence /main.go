package main

import (
	"fmt"
)

/*
Input: nums = [1,3,5,4,7]
              [4,3,2,2,1]
              [2,2,1,1,1]

Output: 2
Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].
 */
func main() {
	//nums := []int{2,2,2,2,2}
	//n := len(nums)
	//cache := []int{1,1,1,1,1}
	//counter := make([]int, n)
	//
	//for i := n-1; i >= 0; i-- {
	//	tempCounter := 0
	//	for j := i + 1; j < n; j++ {
	//		if nums[i] < nums[j] && cache[i] - cache[j] == 1 {
	//			tempCounter += counter[j]
	//		}
	//	}
	//	if tempCounter == 0 {
	//		tempCounter = 1
	//	}
	//	counter[i] = tempCounter
	//}
	//fmt.Println(counter)
	fmt.Println(findNumberOfLIS([]int{1,3,5,4,7}))
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	cache := make([]int, n)

	for i := 0; i < n; i++ {
		length := dfs(nums, i, cache)
		fmt.Println(length)
	}

	counter := make([]int, n)
	for i := n-1; i >= 0; i-- {
		tempCounter := 0
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] && cache[i] - cache[j] == 1 {
				tempCounter += counter[j]
			}
		}
		if tempCounter == 0 {
			tempCounter = 1
		}
		counter[i] = tempCounter
	}

	max := 0
	for i := 0; i < n; i++ {
		if cache[i] > max {
			max = cache[i]
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if cache[i] == max {
			res += counter[i]
		}
	}
	return res
}

func dfs(nums []int, index int, cache []int) int {
	n := len(nums)

	if cache[index] != 0 {
		return cache[index]
	}

	max := 0
	for i := index+1; i < n; i++ {
		if nums[index] >= nums[i] {
			continue
		}
		temp := dfs(nums, i, cache)
		if max < temp {
			max = temp
		}
	}

	cache[index] = max + 1
	return max + 1
}