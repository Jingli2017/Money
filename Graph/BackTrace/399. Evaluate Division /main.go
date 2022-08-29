package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	adj := make(map[string]map[string]float64)
	// build graph
	for i, equ := range equations {
		dividend := equ[0]
		divisor := equ[1]
		if adj[dividend] == nil {
			adj[dividend] = make(map[string]float64)
		}
		adj[dividend][divisor] = values[i]
		if adj[divisor] == nil {
			adj[divisor] = make(map[string]float64)
		}
		adj[divisor][dividend] = 1.0 / values[i]
	}
	// process query
	output := make([]float64, len(queries))
	for i, query := range queries {
		src := query[0]
		dest := query[1]
		if _, ok := adj[src]; !ok {
			output[i] = -1.0
		}else if _, ok := adj[dest]; !ok {
			output[i] = -1.0
		}else if src == dest {
			output[i] = 1.0
		}else{
			visited := make(map[string]bool)
			output[i] = backtrace(adj, visited, 1, src, dest)
		}
	}
	return output
}

func backtrace(adj map[string]map[string]float64, visited map[string]bool, product float64, src, dest string) float64 {
	if _, ok := adj[src][dest]; ok {
		return adj[src][dest] * product
	}
	visited[src] = true
	for nextSrc, nextVal := range adj[src] {
		if visited[nextSrc] { continue }
		res := backtrace(adj, visited, product * nextVal, nextSrc, dest)
		if res != -1.0 {
			return res
		}
	}
	visited[src] = false
	return -1.0
}
