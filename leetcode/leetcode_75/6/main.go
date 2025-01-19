package main

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)

	var builder strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i])
		if i != 0 {
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}