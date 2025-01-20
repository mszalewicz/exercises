package main

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	functionName := "TestReverseWords"
	testCases := []struct {
		text     string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index), func(t *testing.T) {
			if result := reverseWords(test.text); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %s | result: %s", functionName, index, test.expected, result)
			}
		})
	}
}
