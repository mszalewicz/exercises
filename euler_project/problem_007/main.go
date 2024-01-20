package main

import (
	"fmt"
)

func main() {

	const BOUND = 10001
	var primes []int
	i := 1

mainloop:
	for {
		i += 1

		for _, prime := range primes {
			if i%prime == 0 {
				continue mainloop
			}
		}

		primes = append(primes, i)

		if len(primes) == BOUND {
			break mainloop
		}
	}

	fmt.Println(primes[BOUND-1])
}
