package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	functionName := "TestAdd"	

	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{1,5,6},
		{1,5,7},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index), func(t *testing.T){
			if result := add(test.a, test.b); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %d | result: %d", functionName, index, test.expected, result)
			}	
		})
	}
}
