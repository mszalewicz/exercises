package main

// Description:
//
// For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).
// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.
//
// Constraints:
//
//     1 <= str1.length, str2.length <= 1000
//     str1 and str2 consist of English uppercase letters.

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcdOfLen := gcd(len(str1), len(str2))

	return str1[:gcdOfLen]
}

func gcd(a, b int) int {
	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
