package main

import (
	"fmt"
)

func fibo() {
	const limit int = 4000000
	var sum int
	a, b := 0, 1
	for b <= limit {
		next := a + b
		a = b
		b = next
		if a%2 == 0 {
			sum += a
		}
	}
	fmt.Println(sum)
}
