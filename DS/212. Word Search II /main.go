package main

/*
Input: board =
[
	["o","a","a","n"],
	["e","t","a","e"],
	["i","h","k","r"],
	["i","f","l","v"]
						],
words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
*/

type Trie struct {
	children [26]*Trie
	word     string
}

func findWords(board [][]byte, words []string) []string {
	root := &Trie{}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			if node.children[word[i]-'a'] != nil {
				node = node.children[word[i]-'a']
			} else {
				newNode := &Trie{}
				node.children[word[i]-'a'] = newNode
				node = newNode
			}
		}
		node.word = word
	}

	output := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if root.children[board[i][j] - 'a'] != nil {
				backtrace(board, root, i, j, &output)
			}
		}
	}
	return output
}

func backtrace(board [][]byte, node *Trie, i, j int, output *[]string){
	m := len(board)
	n := len(board[0])

	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}

	if board[i][j] == '#' || node.children[board[i][j]-'a'] == nil {
		return
	}

	curr := node.children[board[i][j]-'a']
	if curr.word != "" {
		*output = append(*output, curr.word)
		curr.word = ""
	}

	letter := board[i][j]
	board[i][j] = '#'

	backtrace(board, curr, i+1, j, output)
	backtrace(board, curr, i-1, j, output)
	backtrace(board, curr, i, j+1, output)
	backtrace(board, curr, i, j-1, output)

	board[i][j] = letter
}
