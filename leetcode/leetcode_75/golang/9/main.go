package main

import (
	"fmt"
	"strconv"
)

// Given an array of characters chars, compress it using the following algorithm:
//
// Begin with an empty string s. For each group of consecutive repeating characters in chars:
//
//     If the group's length is 1, append the character to s.
//     Otherwise, append the character followed by the group's length.
//
// The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.
//
// After you are done modifying the input array, return the new length of the array.
//
// You must write an algorithm that uses only constant extra space.

func main() {
	test := []byte{'a', 'a', 'b','b','b','b','b','b','b','b','b','b','b','b','c'}
	result := compress(test)

	fmt.Println(result)

	for i := 0; i < result; i++ {
		fmt.Print(string(test[i]))
	}

	fmt.Println()
}

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	// encode if given byte changes
	howManyWritten := 0
	previous := chars[0]
	counter := 1
	positionToStartWritting := 0

	for i := 1; i < len(chars); i++ {
		if chars[i] != previous {

			if counter == 1 {
				chars[positionToStartWritting] = previous
				previous = chars[i]
				positionToStartWritting++
				howManyWritten++
				counter = 1
				continue
			}

			stringForm := strconv.Itoa(counter)

			chars[positionToStartWritting] = previous
			previous = chars[i]
			positionToStartWritting++
			howManyWritten++

			for _, val := range stringForm {
				chars[positionToStartWritting] = byte(val)
				positionToStartWritting++
				howManyWritten++
			}
			counter = 1
			continue
		}

		counter++
	}

	// write last encoded part
	if counter == 1 {
		chars[positionToStartWritting] = previous
		positionToStartWritting++
		howManyWritten++
	} else {
		stringForm := strconv.Itoa(counter)

		chars[positionToStartWritting] = previous
		positionToStartWritting++
		howManyWritten++

		for _, val := range stringForm {
			chars[positionToStartWritting] = byte(val)
			positionToStartWritting++
			howManyWritten++
		}
	}

	return howManyWritten
}
