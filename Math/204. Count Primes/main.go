package main

import (
	"fmt"
	"math"
)

/*
	Input: n = 10
	Output: 4
	Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 */
func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	notPrime := make([]bool, n)
	fmt.Println(int(math.Sqrt(float64(n))))
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if notPrime[i] {
			continue
		}
		for j := i; i * j < n; j++ {
			notPrime[j * i] = true
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
		}
	}

	return count
}
