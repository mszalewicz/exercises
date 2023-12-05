package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Instruction struct {
	destination int
	source      int
	rangeLength int
}

func traverse(start int, instructionSet *[][]Instruction) int {
	destination := start

	for _, instructions := range *instructionSet {
		for _, instruction := range instructions {
			if destination >= instruction.source && destination <= instruction.source+instruction.rangeLength-1 {
				destination = destination - instruction.source + instruction.destination
				break
			}
		}
	}

	return destination
}

func main() {

	input, _ := os.ReadFile("input/day_05")
	start := time.Now()

	parts := strings.Split(string(input), "\n\n")
	seeds := strings.Fields(parts[0][7:])

	var instructionSet [][]Instruction

	for _, section := range parts[1:] {

		lines := strings.Split(section, "\n")[1:]

		var instructions []Instruction
		var instruction Instruction

		for _, values := range lines {
			numbers := strings.Fields(values)
			destination, _ := strconv.Atoi(numbers[0])
			source, _ := strconv.Atoi(numbers[1])
			rangeLength, _ := strconv.Atoi(numbers[2])

			instruction.destination = destination
			instruction.source = source
			instruction.rangeLength = rangeLength
			instructions = append(instructions, instruction)
		}

		instructionSet = append(instructionSet, instructions)
	}

	var seedSets []int
	result1 := math.MaxInt32

	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		seedSets = append(seedSets, seedInt)
		finalLocation := traverse(seedInt, &instructionSet)

		if finalLocation < result1 {
			result1 = finalLocation
		}
	}

	result2 := math.MaxInt32

	for i := 0; i < len(seedSets)-1; i = i + 2 {
		for j := seedSets[i]; j < seedSets[i]+seedSets[i+1]; j++ {
			finalLocation := traverse(j, &instructionSet)

			if finalLocation < result2 {
				result2 = finalLocation
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println("time " + fmt.Sprint(elapsed))
}
