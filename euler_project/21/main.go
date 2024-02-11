package main

import (
	"fmt"
	"math"
	"slices"
)

func sumSlice(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func main() {
	upperBound := 10000
	sumOfProperDivisors := make([]int, 0, upperBound)
	sumOfAmicableNumbers := 0

	for i := 1; i <= upperBound; i++ {
		properDivisors := []int{}

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				properDivisors = append(properDivisors, j)
				if j != i/j {
					properDivisors = append(properDivisors, i/j)
				}
			}
		}

		sumOfProperDivisors = append(sumOfProperDivisors, sumSlice(properDivisors)+1)
	}

	visited := make([]int, 0)
	sizeOfProperDivisors := len(sumOfProperDivisors)

	for index, value := range sumOfProperDivisors {
		if slices.Contains(visited, index) {
			continue
		}

		if value <= sizeOfProperDivisors {
			amicableCandidate := sumOfProperDivisors[value-1]
			if amicableCandidate == index+1 && index != value-1 {
				sumOfAmicableNumbers += value
				sumOfAmicableNumbers += amicableCandidate
				visited = append(visited, value-1)
			}
		}
	}

	fmt.Println(sumOfAmicableNumbers)
}
