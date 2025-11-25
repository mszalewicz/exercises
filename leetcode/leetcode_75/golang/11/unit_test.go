package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	functionName := "Test"
	testCases := []struct {
		text1    string
		text2    string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"a1c", "ahbg1c", true},
		{"", "ahbg1c", true},
		{"a1c", "", false},
		{"acb", "ahbgdc", false},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := isSubsequence(test.text1, test.text2); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %t | result: %t", functionName, index+1, test.expected, result)
			}
		})
	}
}
