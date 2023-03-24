from typing import List


class Solution:
    """
    Input: nums = [1,5,11,5]
    Output: true
    Explanation: The array can be partitioned as [1, 5, 5] and [11].
    """
    def canPartition(self, nums: List[int]) -> bool:
        total = sum(nums)
        if total % 2 != 0:
            return False
        target = int(total / 2)
        memo = [[-1]*(total+1) for i in range(len(nums))]
        return self.can_partition(len(nums)-1, target, nums, memo)

    def can_partition(self, step, target, nums, memo):
        if step == -1:
            return False
        if target == 0:
            return True
        if memo[step][target] != -1:
            return memo[step][target]
        if nums[step] > target:
            return self.can_partition(step-1, target, nums, memo)
        temp_1 = self.can_partition(step-1, target, nums, memo)
        temp_2 = self.can_partition(step-1, target-nums[step], nums, memo)
        result = temp_1 or temp_2
        memo[step][target] = result
        return result


if __name__ == '__main__':
    S = Solution()
    print(S.canPartition([1, 5, 11, 5]))



