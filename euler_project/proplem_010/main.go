package main

import "fmt"

func main() {
	var primes []int
	result := 0

mainloop:
	for i := 2; i < 2000000; i++ {
		for _, prime := range primes {
			if i%prime == 0 {
				continue mainloop
			}
		}
		result += i
		primes = append(primes, i)
	}
	fmt.Println(result)
}
