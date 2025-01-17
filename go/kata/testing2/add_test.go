package main

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	functionName := "TestAddition"

	testCases := []struct {
		number1  int
		number2  int
		expected int
	}{
		{1, 1, 3},
		{2, 3, 6},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprint(functionName, " case ", index), func(t *testing.T) {
			if result := add(test.number1, test.number2); result != test.expected {
				t.Fatalf("%s number %d failed - expected: %d | result: %d", functionName, index, test.expected, result)
			}
		})
	}
}
