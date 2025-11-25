package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	functionName := "Test"

	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := longestSubarray(test.nums); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %d | result: %d", functionName, index+1, test.expected, result)
			}
		})
	}
}
