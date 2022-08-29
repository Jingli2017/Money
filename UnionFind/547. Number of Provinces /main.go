package _47__Number_of_Provinces_

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && isConnected[i][j] == 1 {
				if find(i, parent) != find(j, parent) {
					union(i, j, parent, size)
				}
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if i == parent[i] {
			count++
		}
	}

	return count
}

func find(p int, parent []int) int {
	root := p
	for root != parent[root] {
		root = parent[root]
	}

	for p != root {
		next := parent[p]
		parent[p] = root
		p = next
	}

	return root
}

func union(i, j int, parent []int, size []int) {
	iparent := find(i, parent)
	jparent := find(j, parent)

	if iparent == jparent {
		return
	}

	if size[iparent] < size[jparent]{
		size[jparent] += size[iparent]
		parent[iparent] = jparent
		size[iparent] = 0
	}else{
		size[iparent] += size[jparent]
		parent[jparent] = iparent
		size[jparent] = 0
	}
}
