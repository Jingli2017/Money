package main

func main() {

}

type node struct {
	level int
	word  string
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	visited := make([]bool, len(wordList))
	q := make([]node, 0)
	q = append(q, node{1, beginWord})

	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		for i, word := range wordList {
			if visited[i] || diff(curr.word, word) > 1 {
				continue
			}

			visited[i] = true
			q = append(q, node{curr.level+1, word})
		}
	}

	return 0
}

func diff(A, B string) int {
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			count++
		}
		if count >= 2 {
			return count
		}
	}

	return count
}
