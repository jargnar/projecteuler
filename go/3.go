package main

import (
	"fmt"
)

func max(arr []int) int {
	var max = arr[0]
	for _, elem := range arr {
		if elem > max {
			max = elem
		}
	}
	return max
}

func largestPrimeFactor() {
	const N int = 600851475143
	factors, _ := primefactors(N)
	fmt.Println(max(factors))
}
