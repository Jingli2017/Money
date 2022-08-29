package _18__Coin_Change_2_

/*
	Input: amount = 5, coins = [1,2,5]
	Output: 4
	Explanation: there are four ways to make up the amount:
	5=5
	5=2+2+1
	5=2+1+1+1
	5=1+1+1+1+1
 */
type combination struct {
	amount int
	step int
}

func change(amount int, coins []int) int {
	dp := make(map[combination]int)
	return impl(amount, 0, coins, dp)
}

func impl(remained int, step int, coins []int, dp map[combination]int) int {
	// base case
	if step == len(coins) {
		if remained == 0 {
			return 1
		}else {
			return 0
		}
	}
	// done before
	if _, ok := dp[combination{remained, step}]; ok {
		return dp[combination{remained, step}]
	}

	total := 0
	coin := coins[step]
	for i := 0; i * coin <= remained; i++ {
		ways := impl(remained - i*coin, step+1, coins, dp)
		total += ways
	}

	dp[combination{remained, step}] = total
	return total
}
