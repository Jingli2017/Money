package _57__Prison_Cells_After_N_Days

/*
	Input: cells = [0,1,0,1,1,0,0,1], n = 7
	Output: [0,0,1,1,0,0,0,0]
	Explanation: The following table summarizes the state of the prison on each day:
	Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
	Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
	Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
	Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
	Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
	Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
	Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
	Day 7: [0, 0, 1, 1, 0, 0, 0, 0]
*/

func prisonAfterNDays(cells []int, n int) []int {
	lists := make([][]int, 0)
	var temp []int
	for i := 0; i < n; i++ {
		temp = make([]int, len(cells))
		for j := 1; j < len(cells)-1; j++ {
			if cells[j-1] == cells[j+1] {
				temp[j] = 1
			}else {
				temp[j] = 0
			}
		}

		if len(lists) != 0 && isEqual(lists[0], temp) {
			k := (n-1) % i
			return lists[k]
		}

		lists = append(lists, temp)
		cells = temp
	}

	return temp
}

func isEqual(A, B []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			return false
		}
	}
	return true
}
