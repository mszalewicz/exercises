package main

import (
	"fmt"
	"testing"
)

func TestMaxAverage(t *testing.T) {
	functionName := "TestMaxAverage"

	testCases := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5.0},
		{[]int{0}, 1, 0.0},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := findMaxAverage(test.nums, test.k); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %f | result: %f", functionName, index+1, test.expected, result)
			}
		})
	}
}
