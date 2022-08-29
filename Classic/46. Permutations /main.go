package main

import "fmt"

/*
	Input: nums = [1,2,3]
	Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func main() {
	fmt.Println(permute([]int{1,2,3}))
}

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	ret := make([][]int, 0)
	ret = append(ret, []int{})

	for _, num := range nums {
		temp := make([][]int, 0)
		for i := 0; i < len(ret); i++ {
			for j := 0; j <= len(ret[i]); j++ {
				temp = append(temp, insert(num, ret[i], j))
			}
		}
		ret = temp
	}

	return ret
}

func insert(num int, arr []int, index int) []int {
	ret := make([]int, len(arr))
	copy(ret, arr)

	if index == len(ret){
		ret = append(ret, num)
		return ret
	}

	ret = append(ret[:index+1], ret[index:]...)
	ret[index] = num

	return ret
}
