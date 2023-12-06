package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumbers(text string) []int {
	var numbers []int
	splitText := strings.Fields(text)[1:]

	for _, numberText := range splitText {
		number, _ := strconv.Atoi(numberText)
		numbers = append(numbers, number)
	}
	return numbers
}

func multiplyWins(wins *[]int) int {
	result := 1
	for _, score := range *wins {
		result *= score
	}
	return result
}

func main() {

	input, _ := os.ReadFile("input/day_06")

	times := getNumbers(strings.Split(string(input), "\n")[0])
	distances := getNumbers(strings.Split(string(input), "\n")[1])
	var wins []int

	for position, time := range times {
		var counter int
		for i := 1; i <= time; i++ {
			if (time-i)*i > distances[position] {
				counter += 1
			}
		}
		if counter != 0 {
			wins = append(wins, counter)
		}
	}

	result := multiplyWins(&wins)

	fmt.Println(result)
}
