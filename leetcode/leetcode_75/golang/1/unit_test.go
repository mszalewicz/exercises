package main

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	functionName := "TestMerge"

	testCases := []struct {
		word1    string
		word2    string
		expected string
	}{
		{"abcd", "123", "a1b2c3d"},
		{"abc", "1234", "a1b2c34"},
		{"", "1234", "1234"},
		{"abcd", "", "abcd"},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index), func(t *testing.T) {
			result := mergeAlternately(test.word1, test.word2)

			if result != test.expected {

				t.Fatalf("%s case number %d failed - expected: %v | result: %v", functionName, index+1, test.expected, result)
			}
		})
	}
}
