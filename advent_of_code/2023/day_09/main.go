package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	// "github.com/mszalewicz/simplemath"
)

func bootstrapStructure(numbers *[][]int, row *[]int, line *[]string) {
	for index, number := range *line {

		convertedNumber, err := strconv.Atoi(number)

		if err != nil {
			log.Fatal(err)
		}

		*row = append(*row, convertedNumber)

		if index == len(*line)-1 {
			*numbers = append(*numbers, *row)
		}
	}
}

func populateStructure(numbers *[][]int) {
	var row []int
	lastRowIndex := 0
	stop := true

	for {
		row = (*numbers)[lastRowIndex]
		var nextRow []int
		for i := 0; i < len(row)-1; i++ {
			// lowerNumber := int(simplemath.Distance(float64(row[i]), float64(row[i+1])))
			lowerNumber := row[i+1] - row[i]
			nextRow = append(nextRow, lowerNumber)
		}
		*numbers = append(*numbers, nextRow)

		for _, value := range nextRow {
			if value != 0 {
				stop = false
			}
		}

		if stop {
			break
		}
		// if nextRow[len(nextRow)-1] == 0 {
		// 	break
		// }

		lastRowIndex += 1
		stop = true
	}
}

func calculateResult(numbers *[][]int) int {
	result := 0
	firstStep := true

	for i := len(*numbers) - 1; i > 0; i-- {

		if firstStep {
			n1 := (*numbers)[i][len((*numbers)[i])-1]
			n2 := (*numbers)[i-1][len((*numbers)[i-1])-1]

			result = n1 + n2

			firstStep = false
		} else {
			result = result + (*numbers)[i-1][len((*numbers)[i-1])-1]
		}

	}

	return result
}

func main() {

	start := time.Now()

	content, err := os.ReadFile("input/day_09")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	result1 := 0

	for _, line := range lines {

		var numbers [][]int
		var row []int

		line := strings.Fields(line)

		bootstrapStructure(&numbers, &row, &line)
		populateStructure(&numbers)
		result1 += calculateResult(&numbers)
	}

	elapsed := time.Since(start)
	fmt.Println(result1)
	fmt.Println(elapsed)
}
