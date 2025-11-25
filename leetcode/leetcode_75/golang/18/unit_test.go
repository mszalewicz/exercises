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
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := largestAltitude(test.nums); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %d | result: %d", functionName, index+1, test.expected, result)
			}
		})
	}
}
