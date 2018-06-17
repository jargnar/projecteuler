package main

import (
	"fmt"
	"math"
)

func nthprime() {
	n := 10001
	lim := int(math.Log(float64(2*n)) * float64(2*n))

	isPrime := make(map[int]bool)
	for i := 2; i <= lim; i++ {
		isPrime[i] = true
	}

	counter := 0

	for i := 2; i <= lim; i++ {
		if isPrime[i] == true {
			counter = counter + 1
			if counter == n {
				fmt.Println(i)
				break
			}
			for j := i; j <= lim; j = j + i {
				isPrime[j] = false
			}
		}
	}
}
