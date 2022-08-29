package _22__Coin_Change

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	return impl(coins, amount, dp)
}

func impl(coins []int, amount int, dp []int) int {
	if amount == 0 { return 0 }
	if amount < 0 { return -1 }
	if dp[amount] != 0 { return dp[amount] }

	min := math.MaxInt32
	for _, coin := range coins {
		tempMin := impl(coins, amount - coin, dp)
		if tempMin < 0 { continue }
		min = getMin(min, tempMin+1)
	}

	if min == math.MaxInt32 {
		min = -1
	}

	dp[amount] = min

	return min
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
