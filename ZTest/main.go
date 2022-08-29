package main

import (
	"fmt"
	"time"
)

type trie struct {
	children [26]int
}

type People struct {
	name string
}

func main() {
	p1 := &People{"jing"}
	p2 := &People{"li"}

	m := make(map[string]*People)
	m["jing"] = p1
	m["li"] = p2

	res := make([]People, 0)
	for _, v := range m {
		go func() {
			fmt.Println(v)
		}()
		res = append(res, *v)
	}
	fmt.Println(res)
	time.Sleep(1000000)
}

func findkth(k int, arr []int, lo, hi int) int {
	if lo == hi {
		return arr[lo]
	}

	val := arr[lo]
	i := lo + 1
	j := hi

	for true {
		for arr[i] < val && i < hi {
			i++
		}

		for arr[j] > val && j > lo {
			j--
		}

		if i >= j {
			break
		}

		swap(arr, i, j)
	}
	swap(arr, lo, j)

	if j == k {
		return arr[k]
	} else if j < k {
		return findkth(k, arr, j+1, hi)
	} else {
		return findkth(k, arr, lo, j-1)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
