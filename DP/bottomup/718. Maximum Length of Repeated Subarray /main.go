package main

import "fmt"

/*
	Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
	Output: 3
	Explanation: The repeated subarray with maximum length is [3,2,1].
 */
func main() {
	fmt.Println(findLength([]int{1,2,3,2,1}, []int{3,2,1,4,7}))
}

func findLength(A []int, B []int) int {
	m := len(A)
	n := len(B)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	max := 0
	index := -1
	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
				if max < dp[i][j] {
					max = dp[i][j]
					index = i
				}
			}
		}
	}

	for i := index; i < index + max; i++ {
		fmt.Println(A[i])
	}
	return max
}
