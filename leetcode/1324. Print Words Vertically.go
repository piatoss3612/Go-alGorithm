package main

import "strings"

// 1324. Print Words Vertically
// Runtime: 1ms
// Memory Usage: 1.94MB
// Time Complexity: O(nm) where n is the number of words and m is the maximum length of a word
// Space Complexity: O(m) where m is the maximum length of a word
func printVertically(s string) []string {
	words := strings.Split(s, " ")

	maxLen := 0

	for _, word := range words {
		if maxLen < len(word) {
			maxLen = len(word)
		}
	}

	ans := make([]string, maxLen)

	for i := 0; i < maxLen; i++ {
		var sb strings.Builder

		for _, word := range words {
			if i < len(word) {
				sb.WriteByte(word[i])
			} else {
				sb.WriteByte(' ')
			}
		}

		ans[i] = strings.TrimRight(sb.String(), " ")
	}

	return ans
}
