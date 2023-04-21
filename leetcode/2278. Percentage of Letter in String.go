package main

import (
	"strings"
)

// 1st solution
// Runtime: 2ms
// Memory Usage: 2MB
// Time Complexity: O(n) -> n = len(s)
// Space Complexity: O(1)
func percentageLetter(s string, letter byte) (ans int) {
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			ans += 1
		}
	}
	ans = (ans * 100) / len(s)
	return
}

// 2nd solution
// Runtime: 2ms
// Memory Usage: 2MB
// Time Complexity: O(n) -> n = len(s)
// Space Complexity: O(1)
func percentageLetter2(s string, letter byte) int {
	cnt := strings.Count(s, string(letter))
	return (cnt * 100) / len(s)
}
