package main

import (
	"fmt"
	"strings"
)

// Description:
//
// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.
//
// Return the merged string.

func main() {
	a := "abcdefg"
	b := "12345"

	fmt.Println(mergeAlternately(a, b))
}

func mergeAlternately(word1 string, word2 string) string {
	maxLength := 0

	if len(word1) > len(word2) {
		maxLength = len(word1)
	} else {
		maxLength = len(word2)
	}

	var builder strings.Builder

	for i := range maxLength {
		if i < len(word1) {
			builder.WriteRune(rune(word1[i]))
		}

		if i < len(word2) {
			builder.WriteRune(rune(word2[i]))
		}
	}

	return builder.String()
}
