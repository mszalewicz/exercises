package main

import (
	"fmt"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	functionName := "TestMoveZeros"
	testCases := []struct {
		numbers  []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{1, 0, 2, 3, 0, 10, 0}, []int{1, 2, 3, 10, 0, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{0, 0, 1, 0, 3, 1, 2}, []int{1, 3, 1, 2, 0, 0, 0}},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {

			moveZeroes(test.numbers)

			for i, val := range test.numbers {
				if val != test.expected[i] {
					t.Fatalf("%s case number %d failed - expected: %v | result: %v", functionName, index+1, test.expected, test.numbers)
				}
			}
		})
	}
}
