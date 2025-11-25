package main

import (
	"fmt"
	"testing"
)

func TestMaxOperations(t *testing.T) {
	functionName := "TestMaxOperations"
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
		{[]int{2, 2, 2, 3, 1, 1, 4, 1}, 4, 2},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := maxOperations(test.nums, test.k); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %d | result: %d", functionName, index+1, test.expected, result)
			}
		})
	}
}
