package main

import "fmt"

type pair struct {
	s1 string
	s2 string
}

func main() {
	studentCoursePair := [][]string{
		{"58", "Software Design"},
		{"58", "Linear Algebra"},
		{"94", "Art History"},
		{"94", "Operating Systems"},
		{"17", "Software Design"},
		{"58", "Mechanics"},
		{"58", "Economics"},
		{"17", "Linear Algebra"},
		{"17", "Political Science"},
		{"94", "Economics"},
		{"25", "Economics"},
	}

	queries := [][]string{
		{"58", "17"},
		{"58", "94"},
		{"58", "25"},
		{"94", "25"},
		{"17", "94"},
		{"17", "25"},
	}

	res := findStudentSharedCourse(studentCoursePair, queries)
	fmt.Println(res)
	fmt.Println(len(res))
}

func findStudentSharedCourse(studentCoursePair [][]string, queries [][]string) map[pair][]string {
	studentToCourse := make(map[string]map[string]bool)
	for _, p := range studentCoursePair {
		student := p[0]
		course := p[1]
		if _, ok := studentToCourse[student]; !ok {
			studentToCourse[student] = make(map[string]bool)
		}
		studentToCourse[student][course] = true
	}

	output := make(map[pair][]string)
	for _, p := range queries{
		s1 := p[0]
		s2 := p[1]
		combination := pair{s1, s2}
		output[combination] = make([]string, 0)
		for course := range studentToCourse[s1] {
			if _, ok := studentToCourse[s2][course]; ok {
				output[combination] = append(output[combination], course)
			}
		}
	}
	return output
}


