package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	functionName := "Test"
	testCases := []struct {
		numbers  []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1,1,0,-3,3}, []int{0,0,9,0,0}},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index), func(t *testing.T) {
			result := productExceptSelf(test.numbers)
			for i, val := range result {
				if val != test.expected[i] {
					t.Fatalf("%s case number %d failed - expected: %v | result: %v", functionName, index, test.expected, result)
				}
			}
		})
	}
}
