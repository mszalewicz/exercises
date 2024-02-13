package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(1)
	y := big.NewInt(1)
	z := big.NewInt(1)

	index := 2

	for {
		if len(y.String()) == 1000 {
			fmt.Println(index)
			break
		}

		index += 1

		z.Set(x)
		x.Set(y)
		y.Add(y, z)
	}
}
