package main

import (
	"fmt"
	"math"
)

func main() {
	const upperBound = 600851475143
	var primes []int
	var factor int

start:
	for i := 2; i < int(math.Sqrt(upperBound)); i++ {

		for _, prime := range primes {
			if i%prime == 0 {
				continue start
			}
		}

		primes = append(primes, i)

		if upperBound%i == 0 {
			factor = i
		}
	}

	fmt.Println(factor)
}
