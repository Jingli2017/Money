package _95__Find_Median_from_Data_Stream

import "container/heap"

type MedianFinder struct {
	lower  *maxHeap
	higher *minHeap
}

func Constructor() MedianFinder {
	maxHeap := &maxHeap{}
	minHeap := &minHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFinder{
		lower:  maxHeap,
		higher: minHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(*this.lower) == 0 || num < (*this.lower)[0] {
		heap.Push(this.lower, num)
	} else {
		heap.Push(this.higher, num)
	}

	if abs(len(*this.lower)-len(*this.higher)) > 1 {
		this.rebalance()
	}
}

func (this *MedianFinder) FindMedian() float64 {
	lower := this.lower
	higher := this.higher

	if len(*lower) == len(*higher) {
		return (float64((*lower)[0]) + float64((*higher)[0])) / 2.0
	}

	if len(*lower) > len(*higher){
		return float64((*lower)[0])
	}else {
		return float64((*higher)[0])
	}
}

func (this *MedianFinder) rebalance() {
	lower := this.lower
	higher := this.higher

	for abs(len(*lower) - len(*higher)) > 1 {
		if len(*lower) > len(*higher){
			heap.Push(higher, heap.Pop(lower))
		}else{
			heap.Push(lower, heap.Pop(higher))
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
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
