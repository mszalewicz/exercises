package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	var resultSet []int

	// Number with n digits can be equal at most n*9^5
	// No 7 digit number can be mapped to 7*9^5 = 413343, as it is number with only 6 digits.
	// Hence, result will be n <= 6, i.e. not bigger than 6 digit number

	for i := 2; i <= 1000000; i++ {
		powersNumber := 0
		digitsText := strconv.Itoa(i)

		for _, digitText := range digitsText {

			digit, err := strconv.Atoi(string(digitText))

			if err != nil {
				log.Fatal(err)
			}

			powersNumber += int(math.Pow(float64(digit), float64(5)))
		}

		if powersNumber == i {
			resultSet = append(resultSet, i)
		}
	}

	result := 0

	for _, value := range resultSet {
		result += value
	}

	fmt.Println(result)
}
