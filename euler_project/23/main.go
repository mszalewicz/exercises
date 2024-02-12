package main

import (
	"fmt"
	"math"
)

func findProperDivisors(number int) []int {
	properDivisors := []int{1}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			properDivisors = append(properDivisors, i)

			secondDivisor := number / i
			if secondDivisor != i {
				properDivisors = append(properDivisors, secondDivisor)
			}
		}
	}

	return properDivisors
}

func isAbundantNumber(properDivisors []int, number int) bool {
	sum := 0

	for _, integer := range properDivisors {
		sum += integer
	}

	return sum > number
}

func main() {
	abundandNumbers := []int{}
	ceiling := 28123

	for i := 1; i <= ceiling; i++ {
		properDivisors := findProperDivisors(i)
		if isAbundantNumber(properDivisors, i) {
			abundandNumbers = append(abundandNumbers, i)
		}
	}

	// Starting as sum of all integers < 24, as 24 is first integer that can be created by sum of two abundant numbers
	sumOfNotAbundantCreatibles := 276
mainLoop:
	for i := 24; i <= ceiling; i++ {
		for _, abundant1 := range abundandNumbers {
			if abundant1 > i {
				break
			}
			for _, abundant2 := range abundandNumbers {
				if abundant2 > i {
					break
				} else if abundant1+abundant2 == i {
					continue mainLoop
				}
			}
		}
		sumOfNotAbundantCreatibles += i
	}

	fmt.Println(sumOfNotAbundantCreatibles)
}
