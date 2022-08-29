package main

import (
	"container/heap"
	"fmt"
)

/*
   Integer[] arr1 = {1,2,3,4};
   Integer[] arr2 = {2,5,6};
   Integer[] arr3 = {2,5,7};
 */
func main() {
	a1 := []int{1,2,3,4}
	a2 := []int{2,5,6}
	a3 := []int{2,5,7}

	i1 := &iterator{arr: a1}
	i2 := &iterator{arr: a2}
	i3 := &iterator{arr: a3}

	s1 := NewStream(i1)
	s2 := NewStream(i2)
	s3 := NewStream(i3)

	res := kDuplicateStreams([]stream{s1, s2, s3}, 2)

	fmt.Println(res)
}

type iterator struct {
	index int
	arr   []int
}

func (i *iterator) hasNext() bool {
	return i.index < len(i.arr)
}

func (i *iterator) getNext() int {
	if i.hasNext() {
		num := i.arr[i.index]
		i.index++
		return num
	}
	return 0
}

type stream struct {
	val int
	ite *iterator
}

func NewStream(ite *iterator) stream {
	val := ite.getNext()
	return stream{ite: ite, val: val}
}

type minHeap []stream

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minHeap) Less(i, j int) bool {
	return m[i].val < m[j].val
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(stream))
}

func (m *minHeap) Pop() interface{} {
	n := len(*m)
	p := (*m)[n-1]
	*m = (*m)[:n-1]
	return p
}

func kDuplicateStreams(streams []stream, k int) []int {
	res := make([]int, 0)
	// push into heap
	h := &minHeap{}
	heap.Init(h)
	for _, s := range streams {
		heap.Push(h, s)
	}
	// pop and de-duplicate
	for len(*h) != 0 {
		currS := heap.Pop(h).(stream)
		currV := currS.val
		count := 1

		for currS.ite.hasNext() {
			nextV := currS.ite.getNext()
			if nextV == currV {
				continue
			}
			currS.val = nextV
			heap.Push(h, currS)
		}
		// search other
		for len(*h) != 0 && (*h)[0].val == currV {
			count++
			nextS := heap.Pop(h).(stream)
			for nextS.ite.hasNext() {
				nextV := nextS.ite.getNext()
				if nextV == currV {
					continue
				}
				nextS.val = nextV
				heap.Push(h, nextS)
			}
		}
		if count >= k {
			res = append(res, currV)
		}
	}
	return res
}
