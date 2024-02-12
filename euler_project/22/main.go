package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func sumUpValues(names []string) int {
	letterValues := map[rune]int{
		'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6,
		'G': 7, 'H': 8, 'I': 9, 'J': 10, 'K': 11, 'L': 12,
		'M': 13, 'N': 14, 'O': 15, 'P': 16, 'Q': 17, 'R': 18,
		'S': 19, 'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24,
		'Y': 25, 'Z': 26,
	}

	result := 0

	for index, name := range names {
		for _, letter := range name {
			result += letterValues[letter] * (index + 1)
		}
	}
	return result
}

func main() {
	inputFile, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	inputText := strings.ReplaceAll(string(inputFile), "\"", "")
	input := strings.Split(inputText, ",")

	slices.Sort(input)
	result := sumUpValues(input)

	fmt.Println(result)
}
