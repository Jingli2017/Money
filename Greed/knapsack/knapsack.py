from typing import List


def knapsack(cap: int, weights: List[int], values: List[int]) -> int:
    weights.insert(0, 0)
    values.insert(0, 0)
    memo = [[-1] * len(weights) for i in range(cap+1)]
    return _knapsack(cap, len(weights) - 1, weights, values, memo)


def _knapsack(cap: int, step: int, weights: List[int], values: List[int], memo) -> int:
    if step == 0 or cap == 0:
        return 0

    if memo[cap][step] != -1:
        return memo[cap][step]

    if weights[step] > cap:
        return _knapsack(cap, step-1, weights, values, memo)

    temp_1 = values[step] + _knapsack(cap-weights[step], step-1, weights, values, memo)
    temp_2 = _knapsack(cap, step-1, weights, values, memo)

    memo[cap][step] = max(temp_1, temp_2)

    return memo[cap][step]


if __name__ == '__main__':
    test_weights = [1, 2, 4, 2, 5]
    test_values = [5, 3, 5, 3, 2]
    print(knapsack(10, test_weights, test_values))

