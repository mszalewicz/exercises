package main

import (
	"fmt"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	functionName := "TestfuncKidsWithCandies"

	testCases := []struct {
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{[]int{1, 2, 3}, 1, []bool{false, true, true}},
		{[]int{1, 2, 3}, 0, []bool{false, false, true}},
		{[]int{}, 1, []bool{}},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index), func(t *testing.T) {
			result := kidsWithCandies(test.candies, test.extraCandies)

			if len(result) != len(test.expected) {
				t.Fatalf(fmt.Sprintf("%s case number %d failed -  expected lenght: %d | result lenght: %d", functionName, index, len(test.expected), len(result)))
			}

			for i, val := range result {
				if val != test.expected[i] {
					t.Fatalf(fmt.Sprintf("%s case number %d failed - index: %d | expected: %t | result: %t", functionName, index, i, test.expected[i], val))
				}
			}
		})
	}

}
