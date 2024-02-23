package main

import (
	"fmt"
	"slices"
)

func addPrimes(primes *[]int) {
	var primeCandidate int
	if len(*primes) == 0 {
		primeCandidate = 2
	} else {
		primeCandidate = (*primes)[len(*primes)-1]
	}

	newPrimesContainer := make([]int, len(*primes), len(*primes)+100)
	copy(newPrimesContainer, *primes)
	*primes = newPrimesContainer

mainLoop:
	for n := 1; n <= 1000; {

		primeCandidate++

		for _, prime := range *primes {
			if primeCandidate%prime == 0 {
				continue mainLoop
			}
		}

		*primes = append(*primes, primeCandidate)
		n++
	}

}

func main() {
	const minimumSInt = -2147483648
	maximizingA := minimumSInt
	maximizingB := minimumSInt
	maxConsecutivePrimes := 0
	primes := []int{}

	addPrimes(&primes)

	for a := -999; a < 100; a++ {
	bLoop:
		for b := -1000; b <= 1000; b++ {
			howManyConsecutivePrimes := 0
			for n := 0; ; n++ {
				maybePrime := (n * n) + (a * n) + b

				if maybePrime > primes[len(primes)-1] {
					addPrimes(&primes)
				}

				if !slices.Contains(primes, maybePrime) {
					if howManyConsecutivePrimes > maxConsecutivePrimes {
						maximizingA = a
						maximizingB = b
						maxConsecutivePrimes = howManyConsecutivePrimes
					}
					continue bLoop
				}

				howManyConsecutivePrimes++
			}
		}
	}

	fmt.Printf("Longest prime chain: %d\nMaximizing coefficients - a: %d, b: %d\nProduct of coefficients: %d\n",
		maxConsecutivePrimes,
		maximizingA,
		maximizingB,
		maximizingA*maximizingB)
}
