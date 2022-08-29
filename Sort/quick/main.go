package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4}
	Sort(test)
	fmt.Println(test)
}

func Sort(arr []int){
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partition(arr, lo, hi)
	sort(arr, lo, j-1)
	sort(arr, j+1, hi)
}
/*
   1 2 3 4
   j i
 */
func partition(arr []int, lo, hi int) int {
	val := arr[lo]
	i := lo+1
	j := hi

	for true{
		for arr[i] < val && i < hi {
			i++
		}

		for arr[j] > val && j > lo {
			j--
		}

		if i >= j { break }
		swap(arr, i, j)
	}
	swap(arr, lo, j)
	return j
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}




