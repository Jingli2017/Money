package main

import "fmt"

/*
	Question 2:
	optionToPick is no longer available, readers can freely pick any option. write a function (or edit the code from question 1) that take in goodEndingPages, badEndingPages, pageOptions, and return all possible good ending that can be reached. Assume that at least 1 good ending can always be reached.
	NOTE: if a bad ending is reached, readers cannot read further. If and infinite loop occures, simply avoid that route.

	goodEndingPages - [9,15, 30]
	badEndingPages - [17, 25]
	pageOption - [ [3,5,11], [13, 15,17] ]
	Good ending(s) route -> 1-2-3-5-6-7-8-9, 1-2-3-11-12-13-15
	bad ending(s) route -> 1-2-3-11-12-13-17 -> no further reading.
 */
func main() {
	good := []int{9,15,30}
	bad := []int{17,25}
	options := [][]int{{3,5,11}, {13,15,17}}
	fmt.Println(goodEndingPath(good, bad, options))
}

func goodEndingPath(goodEndingPages []int, badEndingPages []int, pageOption [][]int) [][]int {
	goodEndHash := make(map[int]bool)
	for _, goodEnding := range goodEndingPages {
		goodEndHash[goodEnding] = true
	}

	badEndingHash := make(map[int]bool)
	for _, badEnding := range badEndingPages {
		badEndingHash[badEnding] = true
	}

	options := make(map[int][]int)
	for _, option := range pageOption {
		page := option[0]
		first := option[1]
		second := option[2]

		options[page] = append(options[page], first)
		options[page] = append(options[page], second)
	}

	output := make([][]int, 0)
	temp := make([]int, 0)
	visited := make(map[int]bool)

	dfs(goodEndHash, badEndingHash, options, &output, &temp, 1, visited)

	return output
}

func dfs(goodEndHash map[int]bool, badEndingHash map[int]bool,
	options map[int][]int, output *[][]int, temp *[]int, page int, visited map[int]bool) {
	// bad ending
	if _, ok := badEndingHash[page]; ok {
		return
	}
	// visited
	if _, ok := visited[page]; ok {
		return
	}
	// add to final result
	if _, ok := goodEndHash[page]; ok {
		duplicate := make([]int, len(*temp))
		copy(duplicate, *temp)
		duplicate = append(duplicate, page)
		*output = append(*output, duplicate)
		return
	}

	// mark visited
	visited[page] = true
	*temp = append(*temp, page)

	// traverse
	if _, ok := options[page]; !ok {
		page++
		dfs(goodEndHash, badEndingHash, options, output, temp, page, visited)
	}else{
		dfs(goodEndHash, badEndingHash, options, output, temp, options[page][0], visited)
		dfs(goodEndHash, badEndingHash, options, output, temp, options[page][1], visited)
	}

	delete(visited, page)
	*temp = (*temp)[:len(*temp)-1]

	return
}
