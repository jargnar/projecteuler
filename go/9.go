package main

import (
	"fmt"
)

func isPythagoreanTriplet(a int, b int, c int) bool {
	if (a*a)+(b*b) == (c * c) {
		return true
	}
	return false
}

// pythagoreanTriplet uses https://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple
func pythagoreanTriplet() {
Loop:
	for m := 0; m <= 1000; m++ {
		for n := 0; n <= 1000; n++ {
			if ((2*m*m)+(2*m*n)) == 1000 && m > n {
				a := m*m - n*n
				b := 2 * m * n
				c := m*m + n*n
				if isPythagoreanTriplet(a, b, c) {
					fmt.Println(a * b * c)
					break Loop
				}
			}
		}
	}
}
