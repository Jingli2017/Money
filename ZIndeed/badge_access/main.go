package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
badge_records = [  ["Martha",   "exit"],
  ["Paul",     "enter"],. 1point3acres.com/bbs
  ["Martha",   "enter"],
  ["Martha",   "exit"],
  ["Jennifer", "enter"],. more info on 1point3acres.com
  ["Paul",     "enter"],. From 1point 3acres bbs
  ["Curtis",   "enter"],
  ["Paul",     "exit"],
  ["Martha",   "enter"],
  ["Martha",   "exit"],
  ["Jennifer", "exit"],
]
find_mismatched_entries(badge_records)
Expected output: ["Paul", "Curtis"], ["Martha"]
*/

func enterOrExitWOBadge(records [][]string) ([]string, []string) {
	log := make(map[string]int)
	for _, record := range records {
		if record[1] == "enter" {
			log[record[0]]++
		} else {
			log[record[0]]--
		}
	}

	missExit := make([]string, 0)
	missEnter := make([]string, 0)

	for name, time := range log {
		if time < 0 {
			missEnter = append(missEnter, name)
		} else if time > 0 {
			missExit = append(missExit, name)
		}
	}

	return missEnter, missExit
}

/*
badge_records = [
     ["Paul", 1355],
     ["Jennifer", 1910]
     ["John", 830]
     ["Paul", 1315]
     ["John", 835]
     ["Paul", 1405]
     ["Paul", 1630]
     ["John", 855],

     ["John", 915]
     ["John", 930]
     ["Jennifer", 1335]
     ["Jennifer", 730]
     ["John", 1630]
     ]
*/
func main() {
	rawLogs := [][]string{
		{"Paul", "1355"},
		{"Jennifer", "1910"},
		{"John", "830"},
		{"Paul", "1315"},
		{"John", "835"},
		{"Paul", "1405"},
		{"Paul", "1630"},
		{"John", "855"},
		{"John", "915"},
		{"John", "930"},
		{"Jennifer", "1335"},
		{"Jennifer", "730"},
		{"John", "1630"},
	}

	res := findFrequentGuys(rawLogs)
	fmt.Println(res)
}

type record struct {
	name string
	time int
}

func findFrequentGuys(rawLog [][]string) map[string][]string {
	records := make([]record, 0)
	rawLogMap := make(map[string][]string)
	for _, log := range rawLog {
		name := log[0]
		time := log[1]
		var convertedTime int
		if len(time) == 3 {
			hour, _ := strconv.Atoi(string(time[0]))
			min, _ := strconv.Atoi(string(time[1:]))
			convertedTime = hour*60 + min
		} else {
			hour, _ := strconv.Atoi(string(time[0:2]))
			min, _ := strconv.Atoi(string(time[2:]))
			convertedTime = hour*60 + min
		}
		records = append(records, record{name, convertedTime})
		rawLogMap[name] = append(rawLogMap[name], time)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].time < records[j].time
	})

	recordMap := make(map[string][]int)
	for _, r := range records {
		recordMap[r.name] = append(recordMap[r.name], r.time)
	}

	output := make(map[string][]string)
	for name, times := range recordMap {
		if len(times) < 3 {
			continue
		}

		diffs := make([]int, 0)
		for i := 1; i < len(times); i++ {
			diffs = append(diffs, times[i]-times[i-1])
		}

		for i := 0; i < len(diffs)-1; i++ {
			if diffs[i]+diffs[i+1] <= 60 {
				output[name] = rawLogMap[name]
				break
			}
		}
	}

	return output
}
