package main

import (
	"math"
	"strings"
)

// Given a string s, reverse only all the vowels in the string and return it.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
//

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	lastVowelFromEnd := math.MaxInt
	result := []rune(s)

	for currentPosition := 0; currentPosition < len(s); currentPosition++ {
		currentValue := s[currentPosition]

		if currentPosition >= lastVowelFromEnd {
			break
		}

		if strings.Contains(vowels, string(currentValue)) {
			for x := lower(len(s)-1, lastVowelFromEnd-1); x >= currentPosition; x-- {
				if strings.Contains(vowels, string(s[x])) {
					result[currentPosition], result[x] = result[x], result[currentPosition]
					lastVowelFromEnd = x
					break
				}
			}
		}
	}

	return string(result)
}

func lower(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
