package main

import "fmt"

// ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
// [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
func main() {
	dict := Constructor()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")
	fmt.Println(dict.Search(".ad"))
}

type index struct {
	char byte
	pos  int
}

type WordDictionary struct {
	normalIndex   map[string]bool
	invertedIndex map[index]map[int]bool // word_id
	wordLens      map[int]bool
	words         map[int]string
	maxLen        int
}

func Constructor() WordDictionary {
	return WordDictionary{
		normalIndex:   make(map[string]bool),
		invertedIndex: make(map[index]map[int]bool),
		wordLens:      map[int]bool{},
		words:         map[int]string{},
		maxLen:        0,
	}
}

func (this *WordDictionary) AddWord(word string) {
	if _, ok := this.normalIndex[word]; ok || len(word) == 0 {
		return
	}

	this.wordLens[len(word)] = true

	if len(word) > this.maxLen {
		this.maxLen = len(word)
	}

	this.normalIndex[word] = true
	id := len(this.normalIndex)
	this.words[id] = word

	for i := 0; i < len(word); i++ {
		tempIndex := index{char: word[i], pos: i}
		if this.invertedIndex[tempIndex] == nil {
			this.invertedIndex[tempIndex] = make(map[int]bool)
		}
		this.invertedIndex[tempIndex][id] = true
	}
}

func (this *WordDictionary) Search(word string) bool {
	if _, ok := this.normalIndex[word]; ok {
		return true
	}

	if len(word) == 0 || !this.wordLens[len(word)]{
		return false
	}

	allDots := true
	lists := make([]map[int]bool, 0)
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			continue
		}
		allDots = false
		if _, ok := this.invertedIndex[index{word[i], i}]; !ok {
			return false
		} else {
			lists = append(lists, this.invertedIndex[index{word[i], i}])
		}
	}

	if allDots {
		return this.wordLens[len(word)]
	}

	base := AllHasIntersection(lists)
	res := false
	for key := range base {
		if len(this.words[key]) == len(word) {
			res = true
		}
	}

	if res {
		this.normalIndex[word] = true
	}

	return res
}

func AllHasIntersection(lists []map[int]bool) map[int]bool {
	res := make(map[int]bool)

	if len(lists) == 1 {
		return lists[0]
	}

	base := lists[0]
	for key := range base {
		found := true
		for i := 1; i < len(lists); i++ {
			if !lists[i][key]{
				found = false
				break
			}
		}
		if found {
			res[key] = true
		}
	}

	return res
}
