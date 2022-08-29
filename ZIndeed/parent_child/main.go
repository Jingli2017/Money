package main

import "fmt"

/*
parentChildPairs = [  (1, 3), (2, 3), (3, 6), (5, 6),
                   (5, 7), (4, 5), (4, 8), (8, 10)  ]
*/

/*
findNodesWithZeroAndOneParents(parentChildPairs) =>
                                  [ [1, 2, 4],    // Individuals with zero parents
                                  [5, 7, 8, 10] // Individuals with exactly one parent ]
*/
func main() {
	edges := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {8, 10}}
	adj, srcs := relation(edges)
	fmt.Println(hasCommonParent(adj, srcs, 3, 8))
	fmt.Println(hasCommonParent(adj, srcs, 5, 8))
	fmt.Println(hasCommonParent(adj, srcs, 6, 8))
	fmt.Println(hasCommonParent(adj, srcs, 1, 3))
}

func relation(edges [][]int) (map[int][]int, []int) {
	set := make(map[int]bool)
	parentToChild := make(map[int][]int)
	for _, edge := range edges {
		parent := edge[0]
		child := edge[1]
		parentToChild[parent] = append(parentToChild[parent], child)
		set[parent] = true
		set[child] = true
	}

	degree := make(map[int]int)
	for _, children := range parentToChild {
		for _, child := range children {
			degree[child]++
		}
	}

	noParent := make([]int, 0)

	for node, _ := range set {
		if degree[node] == 0 {
			noParent = append(noParent, node)
		}
	}

	fmt.Println(noParent)

	return parentToChild, noParent
}

/*
hasCommonAncestor(parentChildPairs, 3, 8) => false
hasCommonAncestor(parentChildPairs, 5, 8) => true
hasCommonAncestor(parentChildPairs, 6, 8) => true
hasCommonAncestor(parentChildPairs, 1, 3) => false
 */
func hasCommonParent(parentToChild map[int][]int, noParent []int, c1, c2 int) bool {
	for _, p := range noParent {
		set := make(map[int]bool)
		q := make([]int, 0)
		q = append(q, p)

		for len(q) != 0 {
			curr := q[0]
			q = q[1:]

			if curr == c1 || curr == c2 {
				set[curr] = true
				if len(set) == 2 {
					return true
				}
			}

			for _, child := range parentToChild[curr] {
				q = append(q, child)
			}
		}
	}

	return false
}
