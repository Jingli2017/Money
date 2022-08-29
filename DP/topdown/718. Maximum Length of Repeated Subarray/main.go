package main

import "fmt"

func main() {
	A := []int{1,2,3,2,1}
	B := []int{3,2,1,4,7}
	max, index := findLength(A, B)
	fmt.Println(max, index)
	fmt.Println(A[index: index+max])
}

/*
	Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
	Output: 3
	Explanation: The repeated subarray with maximum length is [3,2,1].
*/

type node struct {
	i int
	j int
}

func findLength(A []int, B []int) (int, int) {
	cache := make(map[node]int)
	max := 0
	index := 0
	dfs(A, B, 0, 0, cache, &max, &index)
	return max, index
}

func dfs(A, B []int, i, j int, cache map[node]int, max *int, index *int) int {
	if i == len(A) || j == len(B) {
		return 0
	}

	if _, ok := cache[node{i,j}]; ok {
		return cache[node{i,j}]
	}

	temp := 0
	if A[i] == B[j] {
		temp = 1 + dfs(A, B, i+1, j+1, cache, max, index)
		if temp > *max {
			*max = temp
			*index = i
		}
	}
	dfs(A, B, i+1, j, cache, max, index)
	dfs(A, B, i, j+1, cache, max, index)

	cache[node{i, j}] = temp
	return temp
}
