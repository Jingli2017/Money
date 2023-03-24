from typing import List


# https://www.youtube.com/watch?v=H9bfqozjoqs
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [amount + 1] * (amount + 1)
        dp[0] = 0

        for i in range(1, amount + 1):
            for coin in coins:
                if coin <= i:
                    dp[i] = min(dp[i], 1 + dp[i - coin])

        return dp[amount] if dp[amount] != amount + 1 else -1


class Solution2:
    def coinChange(self, coins: List[int], amount: int) -> int:
        memo = [amount + 1] * (amount + 1)
        return self._coinChange(coins, amount, memo)

    def _coinChange(self, coins, amount, memo):
        if amount == 0:
            return 0
        if amount < 0:
            return -1
        if memo[amount] != amount + 1:
            return memo[amount]

        min_coin = amount + 1
        for coin in coins:
            temp = self._coinChange(coins, amount - coin, memo)
            if temp >= 0:
                min_coin = min(min_coin, temp + 1)
        if min_coin == amount + 1:  # no way to make change
            memo[amount] = -1
        else:
            memo[amount] = min_coin
        return memo[amount]


if __name__ == '__main__':
    map = {}
    print(map is None)
