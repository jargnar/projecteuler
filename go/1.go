package main

import (
	"fmt"
)

// finds the summation of arithmetic progression
// upto lim, first element a, and difference a
func sumOfMultiples(a int, lim int) int {
	var n int
	n = (lim - 1) / a
	return (n / 2) * ((2 * a) + (n-1)*a)
}

func main() {
	var sum int
	sum = sumOfMultiples(3, 1000) + sumOfMultiples(5, 1000) - sumOfMultiples(15, 1000)
	fmt.Println(sum)
}
