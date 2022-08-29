package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
	Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
	Output: 2
*/
func main() {
	fmt.Println(networkDelayTime([][]int{{2,1,1}, {2,3,1},{3,4,1}}, 4, 2))
}

type node struct {
	key  int
	cost int
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make(map[int][]node) // construct graph
	for _, time := range times {
		src := time[0]
		dest := time[1]
		cost := time[2]
		adj[src] = append(adj[src], node{key: dest, cost: cost})
	}

	timeToReach := make([]int, n+1)
	for i := 0; i <= n; i++ {
		timeToReach[i] = math.MaxInt32
	}
	timeToReach[k] = 0 // start point

	h := &minHeap{}
	heap.Init(h)
	heap.Push(h, node{k, 0})

	for len(*h) != 0 {
		currNode := heap.Pop(h).(node)
		currKey := currNode.key
		currCost := currNode.cost

		if currCost > timeToReach[currKey] { continue }
		for _, nextNode := range adj[currKey]{
			nextKey := nextNode.key
			nextCost := nextNode.cost
			if currCost + nextCost < timeToReach[nextKey]{
				timeToReach[nextKey] = currCost + nextCost
				heap.Push(h, node{nextKey, currCost + nextCost})
			}
		}
	}

	maxTime := 0
	for i := 1; i <= n; i++ {
		if maxTime < timeToReach[i]{
			maxTime = timeToReach[i]
		}
	}

	if maxTime == math.MaxInt32{
		return -1
	}

	return maxTime
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
