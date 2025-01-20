package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	functionName := "Test"
	testCases := []struct {
		numbers  []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
		{[]int{20, 100, 10, 12, 5, 13}, true},
		{[]int{0, 4, 2, 1, 0, -1, -3}, false},
		{[]int{1, 5, 0, 4, 1, 3}, true},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := increasingTriplet(test.numbers); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %t | result: %t", functionName, index+1, test.expected, result)
			}
		})
	}
}
