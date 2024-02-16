package main

import (
	"fmt"
	"slices"
)

func main() {
	resultCycleLenght := 0
	result := 0

	for denominator := 2; denominator < 1000; denominator++ {

		reminders := []int{}
		reminder := 1
		shiftedNumerator := 1

		for {

			shiftedNumerator = reminder

			for shiftedNumerator < denominator {
				shiftedNumerator *= 10
			}

			reminder = shiftedNumerator % denominator

			if reminder == 0 {
				break
			}

			if slices.Contains(reminders, reminder) {
				cycleBeginningIndex := slices.Index(reminders, reminder)
				cycleLenght := len(reminders[cycleBeginningIndex:])

				if cycleLenght > resultCycleLenght {
					resultCycleLenght = cycleLenght
					result = denominator
				}

				break
			}

			reminders = append(reminders, reminder)
		}
	}

	fmt.Println(result)
	fmt.Println(resultCycleLenght)
}
