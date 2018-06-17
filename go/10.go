package main

import (
	"fmt"
)

func sumOfPrimes() {
	lim := 2000000

	isPrime := make(map[int]bool)
	for i := 2; i <= lim; i++ {
		isPrime[i] = true
	}

	sum := 0

	for i := 2; i <= lim; i++ {
		if isPrime[i] == true {
			sum = sum + i
			for j := i; j <= lim; j = j + i {
				isPrime[j] = false
			}
		}
	}
	fmt.Println(sum)
}
