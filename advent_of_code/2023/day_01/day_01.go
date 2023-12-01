package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func day_01() {

	var numbers []int
	var result int

	readFile, fileScanner := open_file("inputs/day_01.txt")

	for fileScanner.Scan() {

		var (
			first_digit   rune
			second_digit  rune
			number_text   string
			number_result int
		)

		runes := []rune(fileScanner.Text())
		length := len(runes)

		for _, char := range runes {
			if unicode.IsDigit(char) {
				first_digit = char
				break
			}
		}

		for i := length - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				second_digit = runes[i]
				break
			}
		}

		number_text += string(first_digit)
		number_text += string(second_digit)

		number_result, err := strconv.Atoi(number_text)

		if err != nil {
			fmt.Println("Error during string to int conversion. Error:")
			fmt.Println(err)
			os.Exit(1)
		}

		numbers = append(numbers, number_result)
	}

	readFile.Close()

	for _, number := range numbers {
		result += number
	}

	fmt.Println("Result: " + fmt.Sprint(result))

}
