import functools
from typing import List
from queue import PriorityQueue

"""
Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.
"""


@functools.total_ordering
class Node:
    def __init__(self, dist, point):
        self.dist = dist
        self.point = point

    def __lt__(self, other):
        return self.dist < other.dist

    def __eq__(self, other):
        return self.dist == other.dist


class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        pq = PriorityQueue()
        for point in points:
            dist = point[0] * point[0] + point[1] * point[1]
            pq.put(Node(dist, point))
        result = []
        for i in range(k):
            result.append(pq.get().point)
        return result


if __name__ == '__main__':
    nums = [1, 1, 1, 1, 1, 1, 1, 1, 1]
    for i in range(len(nums)):
        print(i)
        i += 10

