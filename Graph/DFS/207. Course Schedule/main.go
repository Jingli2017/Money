package main

func canFinish(n int, pres [][]int) bool {
	visited := make([]int, n)

	adj := make(map[int][]int)
	for _, pre := range pres{
		course := pre[0]
		request := pre[1]
		adj[course] = append(adj[course], request)
	}

	for course, _ := range adj {
		if visited[course] == 2 { continue }
		isCycle := dfs(adj, visited, course)
		if isCycle { return false }
	}

	return true
}

func dfs(adj map[int][]int, visited []int, course int) bool {
	if visited[course] == 2 { return false }
	if visited[course] == 1 { return true }

	visited[course] = 1
	for _, next := range adj[course]{
		isCycle := dfs(adj, visited, next)
		if isCycle {
			return true
		}
	}

	visited[course] = 2
	return false
}
