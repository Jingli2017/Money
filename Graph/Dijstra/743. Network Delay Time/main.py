import sys
from functools import total_ordering
from queue import PriorityQueue
from typing import List


@total_ordering
class Node:
    def __init__(self, num, cost):
        self.num = num
        self.cost = cost

    def __lt__(self, other):
        return self.cost < other.cost

    def __eq__(self, other):
        return self.cost == other.cost

    def __repr__(self):
        return "num = {} and cost = {}".format(self.num, self.cost)


class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        # build adj graph
        adj = {}
        for time in times:
            start = time[0]
            end = time[1]
            cost = time[2]
            if start not in adj:
                adj[start] = []
            adj[start].append(Node(end, cost))

        time_to_reach = [sys.maxsize] * (n + 1)
        time_to_reach[k] = 0

        queue = PriorityQueue()
        queue.put(Node(k, 0))

        while queue.qsize() != 0:
            node = queue.get()
            if node.cost > time_to_reach[node.num]:
                continue
            if adj.get(node.num) is None:
                continue
            for nei in adj.get(node.num):
                if nei.cost + node.cost < time_to_reach[nei.num]:
                    time_to_reach[nei.num] = nei.cost + node.cost
                    queue.put(Node(nei.num, nei.cost + node.cost))

        # start from 1 find the max time
        max_time = -1
        for i in range(1, len(time_to_reach)):
            if time_to_reach[i] > max_time:
                max_time = time_to_reach[i]

        if max_time == sys.maxsize:
            return -1
        else:
            return max_time