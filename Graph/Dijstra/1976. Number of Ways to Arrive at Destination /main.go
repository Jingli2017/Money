package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	n := 5
	roads := [][]int{{0,1,1000000000}, {1,2,1000000000}, {2,3,1000000000}, {3,4,1000000000}}
	fmt.Println(countPaths(n, roads))
}

type node struct {
	key  int
	cost int64
}

func countPaths(n int, roads [][]int) int {
	ways := make([]int, n)
	ways[0] = 1

	adj := make(map[int][]node)
	for _, road := range roads {
		src := road[0]
		dst := road[1]
		cost := int64(road[2])

		adj[src] = append(adj[src], node{dst, cost})
		adj[dst] = append(adj[dst], node{src, cost})
	}

	timeToReach := make([]int64, n)
	for i := 0; i < n; i++ {
		timeToReach[i] = math.MaxInt64
	}
	timeToReach[0] = 0

	h := &minHeap{}
	heap.Init(h)
	heap.Push(h, node{0, 0})

	for len(*h) != 0 {
		curr := heap.Pop(h).(node)

		if curr.cost > timeToReach[curr.key] {
			continue
		}

		for _, next := range adj[curr.key] {
			if next.cost+curr.cost < timeToReach[next.key] {
				ways[next.key] = ways[curr.key]
				timeToReach[next.key] = next.cost + curr.cost
				heap.Push(h, node{next.key, next.cost + curr.cost})
			} else if next.cost+curr.cost == timeToReach[next.key] { // no need to put it into heap again
				ways[next.key] = (ways[next.key] + ways[curr.key]) % 1000000007
			}
		}
	}

	return ways[n-1]
}

type minHeap []node

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minHeap) Less(i, j int) bool {
	return m[i].cost < m[j].cost
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(node))
}

func (m *minHeap) Pop() interface{} {
	n := len(*m)
	p := (*m)[n-1]
	*m = (*m)[:n-1]
	return p
}
