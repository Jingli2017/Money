package main

import "fmt"

// intervals = [[1,3],[2,6],[8,10],[15,18]]
// [[1,6],[8,10],[15,18]]
func main() {
	output := make([][]int, 0)
	intervals := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if intervals[0][0] != 0 {
		output = append(output, []int{0, intervals[0][0]})
	}

	for i := 0; i < len(intervals)-1; i++ {
		output = append(output, []int{intervals[i][1], intervals[i+1][0]})
	}

	fmt.Println(output)
}
