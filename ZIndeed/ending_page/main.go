package main

import "fmt"

/*
Given a book with a finite number of pages. At each page, it can be either empty or contain 2 options, each option represent the page number that readers must jump to to continue reading, if the page does not have any option, simply go to the next page. if the options is available at a page, readers have to pick one. The book has multiple ending pages, if an ending page is reached, no further reading is needed.

Question 1:
Given endingPages, pageOptions, optionToPick, write a function that take in all these arguments and return the ending page. if the pageOption somehow lead to an infinite loop, return -1. Assume that readers can always reach one ending of infinite loop does not happen.

Example: happy path

endingPages - [9,15, 20] -> array of integer
pageOption - [ [3,5,11] ] -> An array that contains multple ingeter array with format of [this page number has options, option 1, option 2]
[3,5,11] means that at page 3, there are 2 options, option 1 is to jump to page 5, option 2 is to jump to page 11.
optionToPick - 1 -> an integer that tells readers which option they have to pick. with optionToPick as 1 and with pageOption [ [3,5,11], [21, 30, 39] ], this means that at page 3 readers have to jump to page 5, at page 21 readers have to jump to page 30.
1-2-3-5-6-7-8-9 -> at page 1 go to 2 because page 1 has no option. do the same for page 2. at page 3 go to page 5 because readers have to pick option 1, continue to page 9 which is an ending, we stop here and return 9 as result

Example: infinite loop

endingPages - [9,15, 20]
pageOption - [ [3,5,11], [7, 1 , 20] ]
optionToPick - 1
1-2-3-5-6-7-1...
return -1 as result because we hit infinite loop here.
*/
func main() {
	endings := []int{9,15,20}
	pageOptions := [][]int{{3,5,11}}
	fmt.Println(endingPage(endings, pageOptions, 1))
}

func endingPage(endings []int, pageOptions [][]int, option int) int {
	endingMap := make(map[int]bool)
	for _, endPage := range endings {
		endingMap[endPage] = true
	}

	pageOptionMap := make(map[int]int)
	for _, pageOption := range pageOptions {
		pageOptionMap[pageOption[0]] = pageOption[option]
	}

	visited := make(map[int]bool)

	p := 1
	for true {
		if _, ok := visited[p]; ok{
			return -1
		}
		if _, ok := endingMap[p]; ok {
			break
		}
		visited[p] = true

		if _, ok := pageOptionMap[p]; ok {
			p = pageOptionMap[p]
		}else{
			p++
		}
	}

	return p
}
