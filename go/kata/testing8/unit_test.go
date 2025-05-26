package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	functionName := "TestAdd"

	testCases := []struct {
		num1 int
		num2 int
		expected int
	}{
		{1,2,5},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index + 1), func(t *testing.T){
			if result := add(test.num1, test.num2); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %d | result: %d", functionName, index+1, test.expected, result)
			}
		})
	}
}