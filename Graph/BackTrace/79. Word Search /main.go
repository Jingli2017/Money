package _9__Word_Search_

/*
	Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
	Output: true
 */
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && backtrace(board, word, i, j, 0){
				return true
			}
		}
	}

	return false
}

func backtrace(board [][]byte, word string, i int, j int, index int) bool {
	if index == len(word) { return true }
	m := len(board)
	n := len(board[0])

	if i < 0 || i >= m || j < 0 || j >= n { return false }
	if board[i][j] != word[index] { return false }

	board[i][j] = '#'
	found := backtrace(board, word, i-1, j, index+1) ||
		backtrace(board, word, i, j-1, index+1) ||
		backtrace(board, word, i+1, j, index+1) ||
		backtrace(board, word, i, j+1, index+1)

	if found {
		return true
	}

	board[i][j] = word[index]
	return false
}


