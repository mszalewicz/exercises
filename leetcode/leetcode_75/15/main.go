package main

// Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
//
// Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

func maxVowels(s string, k int) int {
	if len(s) == 0 || k == 0 {
		return 0
	}

	if len(s) == k {
		numberOfVowels := 0

		for _, character := range s {
			if isVowel(byte(character)) {
				numberOfVowels++
			}
		}

		return numberOfVowels
	}

	numberOfVowels := 0
	maxNumberOfVowels := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			numberOfVowels++
		}

		maxNumberOfVowels = numberOfVowels
	}

	for iteration := range len(s) - k {
		if isVowel(s[iteration]) {
			numberOfVowels--
		}

		if isVowel(s[k + iteration]) {
				numberOfVowels++
		}

		if numberOfVowels > maxNumberOfVowels {
			maxNumberOfVowels = numberOfVowels
		}
	}

	return maxNumberOfVowels
}

func isVowel(character byte) bool {
	return character == 'a' || character == 'e' ||character == 'i' ||character == 'o' ||character == 'u'
}