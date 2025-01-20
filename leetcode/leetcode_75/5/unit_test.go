package main

import (
	"fmt"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	functionName := "TestReverseVowels"
	testCases := []struct{
		text string
		expected string
	}{
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede"},
		{"u'o;B,vKO\"?,.;?fGI 9;`9T`Z,;iqsn4A.:;'H3s8E9. B4TxfOiB'wvM?q'k:Q`J\"E1T7lo e7c!?glI66?JZh\"6 !C,aK! 1J?f9'`SX4Q!Q4XS`'9f?J1 !Ka,C! 6\"hZJ?66Ilg?!c7e ol7T1E\"J`Q:k'q?Mvw'BiOfxT4B .9E8s3H';:.A4nsqi;,Z`To`;o IGf?;.,?\"OKv,B;o'u", "u'o;B,vKO\"?,.;?fGI 9;`9T`Z,;oqsn4o.:;'H3s8i9. B4TxfAEB'wvM?q'k:Q`J\"O1T7li E7c!?glo66?JZh\"6 !C,eK! 1J?f9'`SX4Q!Q4XS`'9f?J1 !KI,C! 6\"hZJ?66alg?!c7a Il7T1e\"J`Q:k'q?Mvw'BoEfxT4B .9i8s3H';:.O4nsqE;,Z`TA`;i IGf?;.,?\"OKv,B;o'u"},
	}

	for index, test := range testCases {
		t.Run(fmt.Sprintf("%s case %d", functionName, index), func(t *testing.T){
			if result := reverseVowels(test.text); result != test.expected {
				t.Fatalf("%s case number %d failed - expected: %s | result: %s", functionName, index, test.expected, result)
			}
		})
	}
}