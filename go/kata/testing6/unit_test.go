package main

import (
	"fmt"
	"testing"
)

func TestAdds(t *testing.T) {
	functionName := "TestAdds"

	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 1, 2},
		{4, 5, 8},
	}

	for number, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, number+1), func(t *testing.T) {
			if result := Add(test.a, test.b); result != test.expected {
				t.Fatalf("%s test case %d failed - result: %d | expected: %d", functionName, number+1, result, test.expected)
			}
		})
	}
}
