package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {

}

func alertNames(keyName []string, keyTime []string) []string {
	nameToTime := make(map[string][]string)
	for i := 0; i < len(keyName); i++ {
		nameToTime[keyName[i]] = append(nameToTime[keyName[i]], keyTime[i])
	}

	output := make([]string, 0)
	for name, times := range nameToTime {
		if len(times) < 3 {
			continue
		}
		if hasAlert(times) {
			output = append(output, name)
		}
	}
	sort.Strings(output)
	return output
}

func hasAlert(times []string) bool {
	mins := make([]int, 0)
	for _, time := range times {
		mins = append(mins, convertToMin(time))
	}

	sort.Ints(mins)

	diffs := make([]int, 0)
	for i := 1; i < len(mins); i++ {
		diffs = append(diffs, mins[i]-mins[i-1])
	}

	for i := 0; i < len(diffs)-1; i++ {
		if diffs[i]+diffs[i+1] < 60 {
			return true
		}
	}

	return false
}

func convertToMin(time string) int {
	splited := strings.Split(time, ":")
	hour, _ := strconv.Atoi(splited[0])
	min, _ := strconv.Atoi(splited[1])
	return hour * 60 + min
}
