package main

import (
	"container/heap"
	"fmt"
)

func main() {
	ints := []int{1,2,3,4,2,3,1,4,2}
	fmt.Println(medianSlidingWindow(ints, 3))
}

func medianSlidingWindow(nums []int, k int) []float64 {
	output := make([]float64, 0)
	minh := &minHeap{}
	maxh := &maxHeap{}
	heap.Init(minh)
	heap.Init(maxh)

	i := 0
	for i < k {
		heap.Push(maxh, nums[i])
		i++
	}
	balance(maxh, minh)
	output = append(output, getMedian(k, maxh, minh))

	i = 0
	for j := k; j < len(nums); j++ {
		if !removeFrontMax(maxh, nums[i]) {
			removeFrontMin(minh, nums[i])
		}

		if abs(len(*maxh)-len(*minh)) > 1 {
			balance(maxh, minh)
		}

		if len(*maxh) == 0 || nums[j] <= (*maxh)[0] {
			heap.Push(maxh, nums[j])
		} else {
			heap.Push(minh, nums[j])
		}

		if abs(len(*maxh)-len(*minh)) > 1 {
			balance(maxh, minh)
		}

		output = append(output, getMedian(k, maxh, minh))
		i++
	}

	return output
}

func balance(maxh *maxHeap, minh *minHeap) {
	for abs(len(*maxh)-len(*minh)) > 1 {
		if len(*maxh) > len(*minh) {
			heap.Push(minh, heap.Pop(maxh))
		} else {
			heap.Push(maxh, heap.Pop(minh))
		}
	}
}

func getMedian(k int, maxh *maxHeap, minh *minHeap) float64 {
	if k%2 != 0 {
		if len(*maxh) > len(*minh) {
			return float64((*maxh)[0])
		} else {
			return float64((*minh)[0])
		}
	} else {
		return (float64((*maxh)[0]) + float64((*minh)[0])) / 2.0
	}
}

func removeFrontMax(maxh *maxHeap, num int) bool {
	n := len(*maxh)
	for i := 0; i < n; i++ {
		p := (*maxh)[i]
		if p == num {
			(*maxh)[i], (*maxh)[n-1] = (*maxh)[n-1], (*maxh)[i]
			*maxh = (*maxh)[:n-1]
			break
		}
	}
	heap.Init(maxh)
	return len(*maxh) != n
}

func removeFrontMin(minh *minHeap, num int) bool {
	n := len(*minh)
	for i := 0; i < n; i++ {
		p := (*minh)[i]
		if p == num {
			(*minh)[i], (*minh)[n-1] = (*minh)[n-1], (*minh)[i]
			*minh = (*minh)[:n-1]
			break
		}
	}
	heap.Init(minh)
	return len(*minh) != n
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() interface{} {
	n := len(*m)
	p := (*m)[n-1]
	*m = (*m)[:n-1]
	return p
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
