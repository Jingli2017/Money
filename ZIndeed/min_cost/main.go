package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
       n1
   e1 /  \ e3
     n2   n3
 e2 /
   n4
*/
func main() {
	n1 := &node{}
	n2 := &node{}
	n3 := &node{}
	n4 := &node{}
	e1 := edge{n2, 1}
	e2 := edge{n4, 2}
	e3 := edge{n3, 5}
	n1.edges = append(n1.edges, e1)
	n1.edges = append(n1.edges, e3)
	n2.edges = append(n2.edges, e2)

	res := findMinCostGraph(n1)
	for _, e := range res {
		fmt.Println(e)
	}
}

type node struct {
	edges []edge
}

type edge struct {
	node *node
	cost int
}

func getMinPath(node *node) []edge {
	min := math.MaxInt32
	temp := make([]edge, 0)
	res := make([]edge, 0)
	dfs(&res, &min, &temp, 0, node)
	return res
}

func dfs(res *[]edge, min *int, temp *[]edge, cost int, node *node) {
	if node == nil {
		return
	}

	if len(node.edges) == 0 {
		if cost < *min {
			*min = cost
			duplicate := make([]edge, 0) // deep copy
			for _, e := range *temp {
				duplicate = append(duplicate, e)
			}
			*res = duplicate
		}
	}

	for _, e := range node.edges {
		*temp = append(*temp, e)
		dfs(res, min, temp, cost+e.cost, e.node)
		*temp = (*temp)[:len(*temp)-1]
	}
}

// dijkstra
func findMinCostGraph(root *node) []edge {
	costToReach := make(map[*node]int)
	costToReach[root] = 0

	h := &minHeap{}
	heap.Init(h)
	h.Push(edge{
		node: root,
		cost: 0,
	})

	list := make([]edge, 0)
	for len(*h) != 0 {
		currEdge := heap.Pop(h).(edge)
		currNode := currEdge.node
		currCost := currEdge.cost

		if _, ok := costToReach[currNode]; ok && currCost > costToReach[currNode]{
			continue
		}

		list = append(list, currEdge)

		if len(currNode.edges) == 0 { // leaf
			return list
		}

		for _, e := range currNode.edges {
			nextNode := e.node
			nextCost := e.cost

			if _, ok := costToReach[nextNode]; !ok || costToReach[nextNode] > nextCost + currCost {
				costToReach[nextNode] = nextCost + currCost
				heap.Push(h, edge{
					node: nextNode,
					cost: nextCost + currCost,
				})
			}
		}
	}

	return nil
}

type minHeap []edge

func(m minHeap) Len() int {
	return len(m)
}

func(m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func(m minHeap) Less(i, j int) bool {
	return m[i].cost < m[j].cost
}

func(m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(edge))
}

func(m *minHeap) Pop() interface{} {
	n := len(*m)
	p := (*m)[n-1]
	*m = (*m)[:n-1]
	return p
}

