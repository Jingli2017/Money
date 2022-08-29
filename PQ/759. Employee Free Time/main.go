package _59__Employee_Free_Time

import "sort"

func main() {

}

type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	intervals := make([]*Interval, 0)
	for i := 0; i < len(schedule); i++ {
		for j := 0; j < len(schedule[i]); j++ {
			intervals = append(intervals, schedule[i][j])
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[i].End
	})

	output := make([]*Interval, 0)
	curr := intervals[0]
	k := 1
	for k < len(intervals) {
		if curr.End < intervals[k].Start {
			output = append(output, &Interval{curr.End, intervals[k].Start})
		} else {
			if curr.End < intervals[k].End {
				curr.End = intervals[k].End
			}
		}
		k++
	}

	return output
}
