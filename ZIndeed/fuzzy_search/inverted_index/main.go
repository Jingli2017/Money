package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
3
java spring javascript
javascript frontend python
python ruby
4
javascript java
java
ruby
python
*/
// inverted index
/*
	javascript -> [java spring javascript, javascript frontend python]
	java -> [java spring javascript]
	ruby -> [python ruby]
	python -> [javascript frontend python, python ruby]
*/
func main() {
	// inverted index
	index := make(map[string][]resume)
	// handle input
	reader := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(read(reader))
	for i := 0; i < n; i++ {
		line := read(reader)
		buildIndex(index, line, i)
	}
	fmt.Println(index)
	// handle query
	m, _ := strconv.Atoi(read(reader))
	for i := 0; i < m; i++ {
		query := read(reader)
		score := getScore(index, query)
		fmt.Println(score)
	}
}

type resume struct {
	id     int
	skills string
	score  int
}

func buildIndex(index map[string][]resume, line string, id int) {
	resume := resume{
		id:     id,
		skills: line,
		score:  0,
	}

	skills := strings.Split(line, " ")
	for _, skill := range skills {
		index[skill] = append(index[skill], resume)
	}
}

// sorted by score and index order 0 -> 1
func getScore(index map[string][]resume, query string) []resume {
	resumeCount := make(map[resume]int)
	skills := strings.Split(query, " ")
	for _, skill := range skills {
		for _, re := range index[skill] {
			if _, ok := resumeCount[re]; ok {
				resumeCount[re]++
			}else{
				resumeCount[re] = 1
			}
		}
	}

	list := make([]resume, 0)
	for re, count := range resumeCount {
		re.score = count
		list = append(list, re)
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].score == list[j].score {
			return list[i].id < list[j].id
		}
		return list[i].score > list[j].score
	})

	return list
}

func read(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	return string(str)
}
