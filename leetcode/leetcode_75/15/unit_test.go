package main

import (
	"fmt"
	"testing"
)

func TestMaxVowels(t *testing.T) {
	functionName := "TestMaxVowel"

	testCases := []struct {
		s        string
		k        int
		expected int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"eetcode", 3, 2},
		{"aeiou", 5, 5},
		{"", 10, 0},
		{"testi", 0, 0},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := maxVowels(test.s, test.k); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %d | result: %d", functionName, index+1, test.expected, result)
			}
		})
	}
}
