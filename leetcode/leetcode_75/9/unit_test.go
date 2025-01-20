package main

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	functionName := "TestCompress"
	testCases := []struct {
		chars              []byte
		expected           []byte
		expectedEncodedLen int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, []byte{'a', '2', 'b', '2', 'c', '3'}, 6},
		{[]byte{'a'}, []byte{'a'}, 1},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, []byte{'a', 'b', '1', '2'}, 4},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index+1), func(t *testing.T) {

			encodedLen := compress(test.chars)

			if encodedLen != test.expectedEncodedLen {
				t.Fatalf("%s case number %d failed - expected encoded lenght: %d | result encoded lenght: %d", functionName, index+1, test.expectedEncodedLen, encodedLen)
			}

			for i := 0; i < encodedLen; i++ {
				if test.chars[i] != test.expected[i] {
					t.Fatalf("%s case number %d failed - expected: %v | result: %v", functionName, index+1, test.expected, test.chars[:encodedLen])
				}
			}
		})
	}
}
