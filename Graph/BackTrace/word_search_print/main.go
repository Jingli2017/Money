package main

import "fmt"

/*
[
	["A","B","C","E"],
	["S","F","C","S"],
	["A","D","E","E"]
]
"ABCB"
*/
func main() {
	board := [][]byte{
		{'c', 'c', 't', 'n', 'a', 'x'},
		{'c', 'c', 'a', 't', 'n', 't'},
		{'a', 'c', 'n', 'n', 't', 't'},
		{'t', 'n', 'i', 'i', 'p', 'p'},
		{'a', 'o', 'o', 'o', 'a', 'a'},
		{'s', 'a', 'a', 'a', 'o', 'o'},
		{'k', 'a', 'i', 'o', 'k', 'i'},
	}

	strs := []string{"catnip", "cccc", "s", "ant", "aoi", "ki", "aaoo", "ooo"}

	for _, str := range strs{
		duplicated := make([][]byte, len(board))
		for i := 0; i < len(board); i++ {
			duplicated[i] = make([]byte, len(board[0]))
			copy(duplicated[i], board[i])
		}
		fmt.Println(exist(duplicated, str))
	}
}
/*
grid1 = [
	['c', 'c', 't', 'n', 'a', 'x'],
	['c', 'c', 'a', 't', 'n', 't'],
	['a', 'c', 'n', 'n', 't', 't'],
	['t', 'n', 'i', 'i', 'p', 'p'],
	['a', 'o', 'o', 'o', 'a', 'a'],
	['s', 'a', 'a', 'a', 'o', 'o'],
	['k', 'a', 'i', 'o', 'k', 'i'],
]
word1 = "catnip"
word2 = "cccc"
word3 = "s"
word4 = "ant"
word5 = "aoi"
word6 = "ki"
word7 = "aaoo"
word8 = "ooo"

grid2 = [['a']]
word9 = "a"

find_word_location(grid1, word1) => [ (1, 1), (1, 2), (1, 3), (2, 3), (3, 3), (3, 4) ]
find_word_location(grid1, word2) =>
[(0, 0), (1, 0), (1, 1), (2, 1)]
OR [(0, 0), (0, 1), (1, 1), (2, 1)]
find_word_location(grid1, word3) => [(5, 0)]
find_word_location(grid1, word4) => [(0, 4), (1, 4), (2, 4)] OR [(0, 4), (1, 4), (1, 5)]
find_word_location(grid1, word5) => [(4, 5), (5, 5), (6, 5)]
find_word_location(grid1, word6) => [(6, 4), (6, 5)]
find_word_location(grid1, word7) => [(5, 2), (5, 3), (5, 4), (5, 5)]
find_word_location(grid1, word8) => [(4, 1), (4, 2), (4, 3)]
find_word_location(grid2, word9) => [(0, 0)]
 */
func exist(board [][]byte, word string) [][]int {
	output := make([][]int, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && backtrace(board, word, i, j, 0, &output) {
				return output
			}
		}
	}

	return nil
}

func backtrace(board [][]byte, word string, i, j, index int, output *[][]int) bool {
	if index == len(word) {
		return true
	}

	m := len(board)
	n := len(board[0])

	if i < 0 || i >= m || j < 0 || j >= n {
		return false
	}

	if board[i][j] != word[index] {
		return false
	}

	board[i][j] = '#'

	*output = append(*output, []int{i, j})

	// only below and right
	found := backtrace(board, word, i+1, j, index+1, output) || backtrace(board, word, i, j+1, index+1, output)

	if found {
		return true
	}

	board[i][j] = word[index]

	*output = (*output)[:len(*output)-1]

	return false
}
