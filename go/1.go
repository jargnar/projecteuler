package main

import (
	"fmt"
)

// finds the summation of arithmetic progression
// with N elements, first element a, and difference d
func sumToN(a int, n int, d int) int {
	return (n / 2) * ((2 * a) + (n-1)*d)
}

// calls sumToN with approppriate numbers
func multipleSum(a int, lim int) int {
	return sumToN(a, lim/a, a)
}

func main() {
	var sum int
	sum = multipleSum(3, 1000) + multipleSum(5, 1000) - multipleSum(15, 1000)
	fmt.Println(sum)
}
