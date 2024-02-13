package main

import (
	"fmt"
	"strconv"
	"time"
)

func isPalindrome(number int) bool {
	numberText := strconv.Itoa(number)
	var lenght int

	switch len(numberText) % 2 {
	case 0:
		lenght = len(numberText) / 2
	default:
		lenght = (len(numberText) - 1) / 2
	}

	for i := 0; i < lenght; i++ {
		if numberText[i] != numberText[len(numberText)-1-i] {
			return false
		}
	}

	return true
}

func main() {

	start := time.Now()
	number := 0
	palindrome := 0

	for x := 999; x > 900; x-- {
		for y := 999; y > 900; y-- {
			number = x * y

			if isPalindrome(number) && number > palindrome {
				palindrome = number
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println(palindrome)
	fmt.Println(elapsed)
}
