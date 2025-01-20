package main

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	functionName := "TestMaxArea"
	testCases := []struct{
		height []int
		expected int
	}{
		{[]int{1,8,6,2,5,4,8,3,7}, 49},
		{[]int{1,1}, 1},
		{[]int{1,0}, 0},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := maxArea(test.height); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %d | result: %d", functionName, index+1, test.expected, result)
			}
		})
	}
}