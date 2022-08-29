package main

import "fmt"

/*
	Input: heights = [2,1,5,6,2,3]
	Output: 10
	Explanation: The above is a histogram where width of each bar is 1.
	The largest rectangle is shown in the red area, which has an area = 10 units.
*/
func main() {
	fmt.Println(largestRectangleArea([]int{1,2,3,4,5}))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	lstack := make([]int, 0)
	farLeft := make([]int, n)
	lstack = append(lstack, 0)
	farLeft[0] = 0

	for i := 1; i < n; i++ {
		if heights[i] > heights[lstack[len(lstack)-1]] {
			farLeft[i] = i
		}else{
			for len(lstack) != 0 && heights[i] <= heights[lstack[len(lstack)-1]] {
				lstack = lstack[:len(lstack)-1]
			}
			if len(lstack) == 0 {
				farLeft[i] = 0
			}else {
				farLeft[i] = lstack[len(lstack)-1] + 1
			}
		}
		lstack = append(lstack, i)
	}

	rstack := make([]int, 0)
	farRight := make([]int, n)
	rstack = append(rstack, n-1)
	farRight[n-1] = n-1

	for i := n-2; i >= 0; i-- {
		if heights[i] > heights[rstack[len(rstack)-1]] {
			farRight[i] = i
		}else{
			for len(rstack) != 0 && heights[i] <= heights[rstack[len(rstack)-1]] {
				rstack = rstack[:len(rstack)-1]
			}
			if len(rstack) == 0 {
				farRight[i] = n-1
			}else {
				farRight[i] = rstack[len(rstack)-1] - 1
			}
		}
		rstack = append(rstack, i)
	}

	fmt.Println(farLeft)
	fmt.Println(farRight)

	max := 0
	for i := 0; i < n; i++ {
		temp := (farRight[i] - farLeft[i] + 1) * heights[i]
		if temp > max {
			max = temp
		}
	}

	return max
}
