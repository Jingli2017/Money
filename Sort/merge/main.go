package main

import "fmt"

func main() {
	test := []int{9, 1}
	Sort(test)
	fmt.Println(test)
}

func Sort(arr []int) {
	aux := make([]int, len(arr))
	sort(arr, aux, 0, len(arr)-1)
}

func sort(arr, aux []int, lo, hi int)  {
	if lo == hi { return }

	mid := (lo + hi) / 2
	sort(arr, aux, lo, mid)
	sort(arr, aux, mid+1, hi)
	merge(arr, aux, lo, mid, hi)
}

func merge(arr, aux []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		aux[i] = arr[i]
	}

	i := lo
	j := mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		}else if j > hi {
			arr[k] = aux[i]
			i++
		}else {
			if aux[i] > aux[j]{
				arr[k] = aux[j]
				j++
			}else{
				arr[k] = aux[i]
				i++
			}
		}
	}
}
