package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func solution_part_a() {

	// On each line, the calibration value can be found by combining the first digit and the
	// last digit (in that order) to form a single two-digit number.

	var numbers []int
	var result int

	readFile, fileScanner := open_file("inputs/day_01_a.txt")

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

	fmt.Println("Result 1: " + fmt.Sprint(result))

}

func solution_part_b() {

	//It looks like some of the digits are actually spelled out with letters: one, two,
	//three, four, five, six, seven, eight, and nine also count as valid "digits".

	var numbers []int
	var result int

	digits := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	readFile, fileScanner := open_file("inputs/day_01_b.txt")

	for fileScanner.Scan() {

		var (
			first_digit   rune
			second_digit  rune
			number_text   string
			number_result int
		)

		runes := []rune(fileScanner.Text())
		length := len(runes)

		// sliding window for parsing words as digits from the front
		var text_window [5]rune

		//first digit
		for _, char := range runes {

			if unicode.IsDigit(char) {
				first_digit = char
				break
			}

			// shift values in text window
			for pos, _ := range text_window[:4] {
				text_window[pos] = text_window[pos+1]
			}

			// add new rune at the end
			text_window[4] = char

			// check for 3, 4 ,5 letter words
			word_3_letters := digits[string(text_window[2:])]

			if word_3_letters != 0 {
				first_digit = word_3_letters
				break
			}

			word_4_letters := digits[string(text_window[1:])]

			if word_4_letters != 0 {
				first_digit = word_4_letters
				break
			}

			word_5_letters := digits[string(text_window[:])]

			if word_5_letters != 0 {
				first_digit = word_5_letters
				break
			}

		}

		// second digit
		for i := length - 1; i >= 0; i-- {

			if unicode.IsDigit(runes[i]) {
				second_digit = runes[i]
				break
			}

			// shift values in text window
			for j := 4; j > 0; j-- {
				text_window[j] = text_window[j-1]
			}

			// add new rune at the end
			text_window[0] = runes[i]

			// check for 3, 4 ,5 letter words
			word_3_letters := digits[string(text_window[:3])]

			if word_3_letters != 0 {
				second_digit = word_3_letters
				break
			}

			word_4_letters := digits[string(text_window[:4])]

			if word_4_letters != 0 {
				second_digit = word_4_letters
				break
			}

			word_5_letters := digits[string(text_window[:])]

			if word_5_letters != 0 {
				second_digit = word_5_letters
				break
			}

		}

		number_text += string(first_digit)
		number_text += string(second_digit)

		number_result, err := strconv.Atoi(number_text)

		// fmt.Println("test")

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

	fmt.Println("Result 2: " + fmt.Sprint(result))

}

func day_01() {

	// To maintain the original (a) solution, duplicate code in solution_part_a and solution_part_b is
	// replicated instead of being stored in separate functions. This method is selected to enhance
	// the clarity and conciseness when reviewing the solution.

	solution_part_a()
	solution_part_b()

}
