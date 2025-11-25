package main

import (
	"fmt"
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	functionName := "TestCanPlaceFlowers"

	testCases := []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 0, 0, 1}, 2, true},
		{[]int{1, 0, 1, 0, 1, 0, 1}, 2, false},
		{[]int{0, 0, 1, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1, 0, 0}, 2, true},
		{[]int{1, 0, 1, 0, 1, 0, 1}, 0, true},
		{[]int{0, 0, 0, 0, 0, 1, 0, 0}, 0, true},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {
			if result := canPlaceFlowers(test.flowerbed, test.n); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %t | result: %t", functionName, index+1, test.expected, result)
			}
		})
	}
}
