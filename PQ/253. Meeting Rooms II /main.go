package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
[[1293,2986],[848,3846],[4284,5907],[4466,4781],[518,2918],[300,5870]]
*/
func main() {
	fmt.Println(minMeetingRooms([][]int{
		{1293, 2986},
		{848, 3846},
		{4284, 5907},
		{4466, 4781},
		{518, 2918},
		{300, 5870},
	}))
}

func minMeetingRooms(meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	h := &minHeap{}
	heap.Init(h)

	for _, meeting := range meetings {
		if len(*h) != 0 && meeting[0] >= (*h)[0] {
			heap.Pop(h)
		}
		heap.Push(h, meeting[1])
	}

	return len(*h)
}

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	n := len(*m)
	p := (*m)[n-1]
	*m = (*m)[:n-1]
	return p
}
