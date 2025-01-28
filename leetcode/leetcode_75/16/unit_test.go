package main

import (
	"fmt"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	functionName := "TestLongestOnes"

	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 1, 0, 0, 0, 1, 1, 1, 1}, 2, 6},
		{[]int{1, 1, 0, 0, 1, 1, 0, 1, 1}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 0, 1}, 4, 4},
		{[]int{1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1}, 9, 32},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := longestOnes(test.nums, test.k); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %d | result: %d", functionName, index+1, test.expected, result)
			}
		})
	}
}
