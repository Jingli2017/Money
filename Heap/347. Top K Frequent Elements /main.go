package main

import (
	"container/heap"
	"fmt"
)

/*
	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]
*/
func main() {
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3}, 2))
}

type point struct {
	freq int
	num  int
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	h := &maxHeap{}
	heap.Init(h)

	for num, freq := range counter {
		heap.Push(h, point{freq, num})
	}

	res := make([]int, 0)
	for k != 0 {
		k--
		res = append(res, heap.Pop(h).(point).num)
	}

	return res
}

type maxHeap []point

func(m maxHeap) Len() int {
	return len(m)
}

func(m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func(m maxHeap) Less(i, j int) bool {
	return m[i].freq > m[j].freq
}

func(m *maxHeap) Push(x interface{}){
	*m = append(*m, x.(point))
}

func(m *maxHeap) Pop() interface{} {
	n := len(*m)
	p := (*m)[n-1]
	*m = (*m)[:n-1]
	return p
}