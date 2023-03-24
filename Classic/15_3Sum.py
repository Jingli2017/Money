from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        result = []
        for i in range(len(nums)-2):
            if nums[i] > 0:
                return result
            front = i+1
            back = len(nums)-1
            target = -nums[i]
            while front > back:
                if nums[front] + nums[back] > target:
                    back -= 1
                elif nums[front] + nums[back] < target:
                    front += 1
                else:
                    temp = [nums[i], nums[front], nums[back]]
                    result.append(temp)
                    while front < back and nums[front] == temp[0]:
                        front += 1
                    while front < back and nums[back] == temp[1]:
                        back -= 1
            while i + 1 < len(nums)-2 and nums[i+1] == nums[i]:
                i += 1
        return result
