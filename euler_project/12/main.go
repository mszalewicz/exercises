package main

import (
	"fmt"
	"math"
)

func main() {
	result := 1
	nextNaturalNumber := result + 1

	for {
		divisors := []int{1, result}

		for i := 2; i < int(math.Ceil(math.Sqrt(float64(result)))); i++ {
			if result%i == 0 {
				if result/i == i {
					divisors = append(divisors, i)
				} else {
					divisors = append(divisors, i)
					divisors = append(divisors, result/i)
				}
			}
		}

		if len(divisors) > 500 {
			break
		}

		result += nextNaturalNumber
		nextNaturalNumber += 1
	}

	fmt.Println(result)
}
