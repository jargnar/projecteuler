package main

import (
	"fmt"
)

func sumOfSquares(n int) int {
	// returns sum of squares of numbers up to n
	return (n * (n + 1) * (2*n + 1)) / 6
}

func squareOfSum(n int) int {
	// returns square of sum of numbers up to n
	sum := ((n / 2) * (2 + (n - 1)))
	return sum * sum
}

func sumSquareDiff() {
	fmt.Println(squareOfSum(100) - sumOfSquares(100))
}
