package main

import (
	"fmt"
	"math/big"
)

func lcm2(a *big.Int, b *big.Int) *big.Int {
	var z big.Int
	return z.Mul(z.Div(a, z.GCD(nil, nil, a, b)), b)
}

func lcm() {
	var z = big.NewInt(1)
	for i := 1; i <= 20; i++ {
		z.Set(lcm2(big.NewInt(int64(i)), z))
	}

	fmt.Println(z)
}
