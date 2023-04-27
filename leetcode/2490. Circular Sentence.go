package main

import "strings"

// 1st solution - split the sentence into words
// Runtime: 3ms
// Memory Usage: 2.3 MB
// Time complexity: O(n) -> n is the number of words
// Space complexity: O(n)
func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	n := len(words)

	for i := 0; i < n; i++ {
		if !(words[i][len(words[i])-1] == words[(i+1)%n][0]) {
			return false
		}
	}
	return true
}

// 2nd solution
// Runtime: 3ms
// Memory Usage: 2.2 MB
// Time complexity: O(n) -> n is the length of the sentence
// Space complexity: O(1)
func isCircularSentence2(sentence string) bool {
	n := len(sentence)

	if sentence[n-1] != sentence[0] {
		return false
	}

	for i := 1; i < n; i++ {
		if sentence[i] == ' ' && !(sentence[i-1] == sentence[i+1]) {
			return false
		}
	}
	return true
}
