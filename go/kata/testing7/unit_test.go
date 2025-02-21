package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	testName := "TestAdd"

	testCases := []struct {
		a int
		b int
		expected int
	}{
		{1,1,2},
		{1,1,3},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", testName, index+1), func(t *testing.T){
			if result := Add(test.a, test.b); result != test.expected {
				t.Fatalf("Test case %d failed -  expected: %d | result: %d", index+1, test.expected, result)
			}
		})
	}
}
