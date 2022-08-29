package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	counts := []string{
		"POINT,333858038",
		"NOT,4522732626",
		"INTO,1144226142",
		"ON,4594521081",
		"FOR,6545282031",
		"NOW,679337516",
		"ONE,2148983086",
		"BEHAVIOR,104177552",
		"WAITS,2911079",
		"PEOPLE,658716166",
		"HI,15453893",
		"FORM,352032932",
		"OF,30966074232",
		"THROUGH,647091198",
		"BETWEEN,744064796",
		"FOUR,262968583",
		"LEFT,306802162",
		"OFF,302535533",
		"FROM,3469207674",
		"NO,1400645478",
		"FORMS,136468034",
		"A,45916054218",
	}

	wordList := make([]string, 0)
	wordClicks := make(map[string]int)
	for _, count := range counts {
		split := strings.Split(count, ",")
		word := split[0]
		clicks, _ := strconv.Atoi(split[1])
		wordList = append(wordList, word)
		wordClicks[word] += clicks
	}

	fmt.Println(wordList)
	/*
	   NO -> NOT -> INTO -> POINT
	   ladder('NO', 'POINT', index, wordlist) -> ['NO', 'NOT', 'INTO', 'POINT']
	   ladder('OF', 'FORMS', index, wordlist) -> [ 'OF', 'FOR', 'FROM', 'FORMS ]
	   ladder('ON', 'LEFT', index, wordlist) -> None
	*/
	res := ladder("ON", "LEFT", wordList, wordClicks)
	fmt.Println()
	fmt.Println(res)
}

func ladder(src, dst string, wordList []string, wordClick map[string]int) []string {
	q := make([]string, 0)
	q = append(q, src)

	adj := make(map[string][]string)
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		for _, word := range wordList {
			if canAdd(curr, word) {
				adj[curr] = append(adj[curr], word)
				q = append(q, word)
			}
		}
	}

	fmt.Printf("%#v", adj)
	temp := make([]string, 0)
	res := make([][]string, 0)
	dfs(adj, src, dst, temp, &res)

	fmt.Println()
	fmt.Println(res)

	if len(res) == 0 {
		return nil
	}

	max := 0
	maxIndex := 0

	for i, completed := range res {
		tempSum := 0
		for _, word := range completed {
			tempSum += wordClick[word]
		}
		if tempSum > max{
			max = tempSum
			maxIndex = i
		}
	}

	return res[maxIndex]
}

func dfs(adj map[string][]string, src, dst string, temp []string, res *[][]string) {
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

func canAdd(A, B string) bool {
	if len(B)-len(A) != 1 {
		return false
	}

	counter := [26]int{}
	for i := 0; i < len(B); i++ {
		counter[B[i]-'A']++
	}

	for i := 0; i < len(A); i++ {
		counter[A[i]-'A']--
		if counter[A[i]-'A'] < 0 {
			return false
		}
	}

	return true
}
