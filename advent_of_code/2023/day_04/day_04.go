package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func arrayToInt(numbersString *[]string) []int {
	var trueNumbers []int
	for _, text := range *numbersString {
		conversion, _ := strconv.Atoi(text)
		trueNumbers = append(trueNumbers, conversion)
	}
	return trueNumbers
}

func sum(numbers *[]int) int {
	counter := 0
	for _, number := range *numbers {
		counter += number
	}
	return counter
}

func main() {
	content, _ := os.ReadFile("input/day_04")
	start := time.Now()
	cards := strings.Split(string(content), "\n")
	result1 := 0
	var instancesOfGivenCard []int

	for range cards {
		instancesOfGivenCard = append(instancesOfGivenCard, 1)
	}

	for _, card := range cards {
		score := 1
		anyHits := false
		hitCounter := 0
		cardID, _ := strconv.Atoi(strings.Fields(strings.Split(card, ": ")[0][5:])[0])
		numbers := strings.Split(strings.Split(card, ": ")[1], " | ")
		winningNumbers := strings.Fields(numbers[0])
		winningNumbersInt := arrayToInt(&winningNumbers)
		haveNumbers := strings.Fields(numbers[1])
		haveNumbersInt := arrayToInt(&haveNumbers)

		for _, number := range haveNumbersInt {
			if slices.Contains(winningNumbersInt, number) {
				score = score << 1
				anyHits = true
				hitCounter += 1
			}
		}
		if anyHits {
			result1 += score >> 1
			for i := 1; i <= hitCounter; i++ {
				instancesOfGivenCard[cardID+i-1] += instancesOfGivenCard[cardID-1]
			}
		}
	}

	result2 := sum(&instancesOfGivenCard)

	elapsed := time.Since(start)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println("time " + fmt.Sprint(elapsed))

}
