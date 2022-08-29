package main

import "fmt"

/*
	https://leetcode.com/company/indeed/discuss/2105294/Indeed-or-SDE-or-Karat

	Expected output (unordered):
	[ "A": ["K", "H", "I"],
	"E:" ["H", "L", "I"],
	"J": ["M"] ]
*/
func main() {
	edges := [][]string{
		{"A", "B"},
		{"A", "C"},
		{"B", "K"},
		{"C", "K"},
		{"E", "L"},
		{"F", "G"},
		{"J", "M"},
		{"E", "F"},
		{"G", "H"},
		{"G", "I"},
		{"C", "G"},
	}

	fmt.Println(toTheEnd(edges))
}

func toTheEnd(connections [][]string) map[string]map[string]bool {
	adj := make(map[string][]string)
	degree := make(map[string]int)
	nodes := make(map[string]bool)
	for _, connection := range connections {
		src := connection[0]
		dst := connection[1]
		adj[src] = append(adj[src], dst)
		nodes[src] = true
		nodes[dst] = true
	}

	for _, dsts := range adj {
		for _, dst := range dsts {
			degree[dst]++
		}
	}
	// find source, which does not have edge point to it
	source := make([]string, 0)
	for node, _ := range nodes {
		if _, ok := degree[node]; !ok {
			source = append(source, node)
		}
	}
	fmt.Println(source)
	res := make(map[string]map[string]bool)

	// BFS
	for _, src := range source {
		res[src] = make(map[string]bool)
		q := make([]string, 0)
		q = append(q, src)
		for len(q) != 0 {
			curr := q[0]
			q = q[1:]

			if _, ok := adj[curr]; !ok {
				res[src][curr] = true
				continue
			}

			for _, next := range adj[curr] {
				q = append(q, next)
			}
		}
	}

	return res
}
