package main

import "fmt"

func reverseInt(n int) int {
	var reversed int
	for n > 0 {
		rem := n % 10
		reversed = (reversed * 10) + rem
		n /= 10
	}
	return reversed
}

func palindromeProduct() {
	var product int

Loop:
	for j := 999; j > 900; j-- {
		for i := j; i > 900; i-- {
			product = i * j
			if product == reverseInt(product) {
				fmt.Println(product)
				break Loop
			}
		}
	}
}
