package main

import (
	"fmt"
	"slices"
)

func factorial(number int) int {
	result := 1
	for i := range number {
		result *= i + 1
	}
	return result
}

// Heap's Algorithm for generating permutations:
func generatePermutations(digits []rune, size int, allPermutations *[]string) {
	if size == 1 {
		*allPermutations = append(*allPermutations, string(digits))
		return
	}

	for i := range size {
		generatePermutations(digits, size-1, allPermutations)
		if size%2 == 1 {
			digits[0], digits[size-1] = digits[size-1], digits[0]
		} else {
			digits[i], digits[size-1] = digits[size-1], digits[i]
		}
	}
}

func main() {
	digitsText := "0123456789"
	digits := []rune(digitsText)
	lenght := len(digits)

	numberOfPossiblePermutations := factorial(lenght)
	allPermutations := make([]string, 0, numberOfPossiblePermutations)
	generatePermutations(digits, lenght, &allPermutations)
	slices.Sort(allPermutations)

	fmt.Println(allPermutations[999999])
}
