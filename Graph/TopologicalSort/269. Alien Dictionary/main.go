package main

import (
	"fmt"
	"strings"
)

func main() {
	r1 := rule{}
	r2 := rule{}

	fmt.Println(r1 == r2)
}

type rule struct {
	src byte
	dst byte
}

func alienOrder(words []string) string {
	charSet := make(map[byte]bool)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			charSet[word[i]] = true
		}
	}

	ruleSet := make(map[rule]bool)
	emptyRule := rule{}
	for i := 0; i < len(words)-1; i++ {
		tempRule := getRule(words[i], words[i+1])
		if tempRule == emptyRule && len(words[i]) > len(words[i+1]) { // case of [abc ab]
			return ""
		}
		ruleSet[tempRule] = true
	}

	lowerToHighers := make(map[byte][]byte)
	degree := make(map[byte]int)
	for r, _ := range ruleSet {
		lower := r.src
		higher := r.dst
		lowerToHighers[lower] = append(lowerToHighers[lower], higher)
		degree[higher]++
	}

	q := make([]byte, 0)
	for char := range charSet{
		if _, ok := degree[char]; !ok {
			q = append(q, char)
		}

	}

	var b strings.Builder
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		b.WriteString(string(curr))
		for _, higher := range lowerToHighers[curr] {
			degree[higher]--
			if degree[higher] == 0 {
				q = append(q, higher)
			}
		}
	}

	return b.String()
}

func getRule(A, B string) rule {
	var n int
	if len(A) > len(B) {
		n = len(B)
	}else{
		n = len(A)
	}

	for i := 0; i < n; i++ {
		if A[i] != B[i]{
			newRule := rule{A[i], B[i]}
			return newRule
		}
	}

	return rule{}
}