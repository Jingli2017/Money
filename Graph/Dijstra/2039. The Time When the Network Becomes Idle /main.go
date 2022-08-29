package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(networkBecomesIdle([][]int{{0,1},{1,2}}, []int{0, 2, 1}))
}

type node struct {
	key int
	cost int
}

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	adj := make(map[int][]node)
	for _, edge := range edges{
		src := edge[0]
		dest := edge[1]

		adj[src] = append(adj[src], node{dest, 1})
	}

	timeToReach := make([]int, n)
	for i := 0; i < n; i++ {
		timeToReach[i] = math.MaxInt32
	}
	timeToReach[0] = 0 //master

	h := &minHeap{}
	heap.Init(h)
	heap.Push(h, node{0, 0})

	for len(*h) != 0 {
		curr := heap.Pop(h).(node)

		if curr.cost > timeToReach[curr.key] {
			continue
		}

		for _, next := range adj[curr.key]{
			nextKey := next.key
			nextCost := next.cost

			if nextCost + curr.cost < timeToReach[nextKey]{
				timeToReach[nextKey] = nextCost + curr.cost
				heap.Push(h, node{nextKey, nextCost + curr.cost})
			}
		}
	}

	max := 0
	for i := 1; i < n; i++ {
		singleTime := timeToReach[i]
		extraPayload := (2 * singleTime - 1) / patience[i]
		lastout := extraPayload * patience[i]
		lastin := lastout + 2 * singleTime
		if lastin > max {
			max = lastin
		}
	}

	return max+1
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
