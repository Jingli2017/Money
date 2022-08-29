package main

import "fmt"

/*
	Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
	Output: [0,2,1,3]
 */
/*
   3 -> 1 -> 0
      \ 2/
0: [1, 2]
1: [3]
2: [3]

degree[0 1 2 3]
       0 1 1 2

0, 1, 2, 3
 */
func main() {
	fmt.Println(findOrder(4, [][]int{{1,0},{2,0},{3,1},{3,2}}))
}

func findOrder(n int, pres [][]int) []int {
	dependencyToCourse := make(map[int][]int)
	for _, pre := range pres {
		course := pre[0]
		dependency := pre[1]
		dependencyToCourse[dependency] = append(dependencyToCourse[dependency], course)
	}

	degrees := make([]int, n)
	for _, courses := range dependencyToCourse {
		for _, course := range courses {
			degrees[course]++
		}
	}

	q := make([]int, 0)
	for course, degree := range degrees {
		if degree == 0 { // this course has no prerequisite
			q = append(q, course)
		}
	}

	output := make([]int, 0)
	for len(q) != 0 {
		dependency := q[0]
		q = q[1:]
		output = append(output, dependency)

		for _, course := range dependencyToCourse[dependency] {
			degrees[course]--
			if degrees[course] == 0{
				q = append(q, course)
			}
		}
	}

	if len(output) != n {
		return nil
	}

	return output
}
