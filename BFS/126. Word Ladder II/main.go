package main

import "fmt"

/*
"hit"
"cog"
["hot","dot","dog","lot","log","cog"]
 */
func main() {
	res := findLadders("hit", "cog", []string{"hot","dot","dog","lot","log","cog"})
	fmt.Println()
	fmt.Println(res)
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	q := make([]string, 0)
	q = append(q, beginWord)

	visited := make(map[string]int)
	visited[beginWord] = 1

	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	adj := make(map[string][]string)
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		neighbours := findNeighbours(curr, wordSet)
		for _, word := range neighbours {
			if visited[word] == 0 {
				visited[word] = visited[curr] + 1
				adj[curr] = append(adj[curr], word)
				if word == endWord {
					continue
				}
				q = append(q, word)
			}else if visited[word] == visited[curr] + 1 {
				adj[curr] = append(adj[curr], word)
			}
		}
	}

	temp := make([]string, 0)
	res := make([][]string, 0)
	dfs(adj, beginWord, endWord, temp, &res)

	return res
}

func dfs(adj map[string][]string, src, dst string, temp []string, res *[][]string){
	temp = append(temp, src)
	if src == dst {
		completed := make([]string, len(temp))
		copy(completed, temp)
		*res = append(*res, completed)
		return
	}

	for _, next := range adj[src] {
		dfs(adj, next, dst, temp, res)
	}
}

func findNeighbours(curr string, wordSet map[string]bool) []string {
	res := make([]string, 0)
	currArr := []rune(curr)
	var char rune
	for char = 'a'; char <= 'z'; char++ {
		for i := 0; i < len(curr); i++ {
			if currArr[i] == char {
				continue
			}
			origin := currArr[i]
			currArr[i] = char
			if _, ok := wordSet[string(currArr)]; ok {
				res = append(res, string(currArr))
			}
			currArr[i] = origin
		}
	}
	return res
}