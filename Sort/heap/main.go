package main

import "fmt"

func main() {
	test := []int{1, 5, 7, 3}
	h := &heap{}
	h.Sort(test)
	fmt.Println(test)
}

type heap []int

func (h *heap) Sort(nums []int) {
	n := len(nums)
	for k := n/2; k >= 1; k-- { // heap construction
		sink(nums, k, n)
	}

	k := n
	for k > 1{
		swap(nums, 1, k)
		k--
		sink(nums, 1, k)
	}
}

func sink(nums []int, k int, n int) {
	for 2*k <= n {
		j := 2 * k
		if j < n && less(nums, j, j+1) {
			j++
		}
		if less(nums, j, k) {
			break
		}
		swap(nums, j, k)
		k = j
	}
}

func less(nums []int, i, j int) bool {
	return nums[i-1] < nums[j-1]
}

func swap(nums []int, i, j int) {
	temp := nums[i-1]
	nums[i-1] = nums[j-1]
	nums[j-1] = temp
}
