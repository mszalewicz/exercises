package main

import (
	"fmt"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	functionName := "TestGcdOfStrings"

	testCases := []struct {
		str1     string
		str2     string
		expected string
	}{
		{"ABB", "ABBABB", "ABB"},
		{"ABBC", "ABBCABBC", "ABBC"},
		{"ABB", "LEED", ""},
		{"AAA", "A", "A"},
		{"A", "AAAAAAAAA", "A"},
		{"AA", "AAAAAAAA", "AA"},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case number %d: expected = ", functionName, index), func(t *testing.T) {
			if result := gcdOfStrings(test.str1, test.str2); result != test.expected {
				t.Fatalf("%s case %d - expected: %s | result: %s", functionName, index, test.expected, result)
			}
		})
	}
}
