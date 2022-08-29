package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	rawn, _, _ := reader.ReadLine()
	n, _ := strconv.Atoi(string(rawn))
	lines := make([]string, n)
	for i := 0; i < n; i++ {
		str, _, _ := reader.ReadLine()
		lines[i] = string(str)
	}

	rawm, _, _ := reader.ReadLine()
	m, _ := strconv.Atoi(string(rawm))
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		str, _, _ := reader.ReadLine()
		queries[i] = string(str)
	}

	jobs := make([]job, 0)
	for id, line := range lines {
		languages := strings.Split(line, " ")
		counter := make(map[string]int)
		for _, lan := range languages {
			counter[lan]++
		}
		jobs = append(jobs, job{id: id, count: counter})
	}
	fmt.Println(jobs)
	fmt.Println(queries)

	for _, query := range queries {
		selectedJobs := search(jobs, query)
		fmt.Println(selectedJobs)
	}
}

type job struct {
	id    int
	count map[string]int
	score int
}

func search(jobs []job, query string) []job {
	selectedJobs := make([]job, 0)
	for _, j := range jobs {
		queries := strings.Split(query, " ")
		for _, q := range queries {
			if _, ok := j.count[q]; ok {
				j.score += j.count[q]
			}
		}
		if j.score != 0 {
			selectedJobs = append(selectedJobs, j)
		}
	}
	return selectedJobs
}
