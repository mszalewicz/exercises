package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	functionName := "Test"

	testCases := []struct {
		number1  int
		number2  int
		expected int
	}{
		{1, 1, 2},
	}


	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index ), func(t *testing.T){
			if result := add(test.number1, test.number2); result != test.expected {
			t.Fatalf("%s case number %d failed: expected: %d | result: %d", functionName, index, test.expected, result)
			}
		})
	}

}
