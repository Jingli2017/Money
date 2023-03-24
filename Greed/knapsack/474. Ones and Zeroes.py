from typing import List


class Solution:
    def findMaxForm(self, strs: List[str], m: int, n: int) -> int:
        memo = [[[-1]*(n+1) for i in range(m+1)] for j in range(len(strs))]
        return self._findMaxForm(strs, m, n, 0, memo)

    def _findMaxForm(self, strs, m, n, step, memo):
        if step == len(strs) or (m == 0 and n == 0):
            return 0
        if memo[step][m][n] != -1:
            return memo[step][m][n] != -1
        c_m, c_n = getNumberOfZeroAndOne(strs[step])
        if c_m > m or c_n > n:
            return self._findMaxForm(strs, m, n, step + 1, memo)
        temp_1 = self._findMaxForm(strs, m, n, step + 1, memo)
        temp_2 = 1 + self._findMaxForm(strs, m - c_m, n - c_n, step + 1, memo)
        memo[step][m][n] = max(temp_1, temp_2)
        return memo[step][m][n]


def getNumberOfZeroAndOne(binary):
    n_0 = 0
    n_1 = 0
    for char in binary:
        if char == '0':
            n_0 += 1
        else:
            n_1 += 1
    return n_0, n_1


if __name__ == '__main__':
    map = {}
    map[(1, 1, 1)] = 2
    print(map.get((1, 1, 1)))