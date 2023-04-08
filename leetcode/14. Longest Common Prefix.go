package main

import (
	"strings"
)

// 1st solution (Horizontal Scanning)
// Runtime: 3ms
// Memory: 2.4MB
// Time Complexity: O(S) -> S is the sum of all characters in strs slice
// Space Complexity: O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
	}

	return prefix
}

func commonPrefix(a, b string) string {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	idx := -1

	for i := 0; i < minLen; i++ {
		if a[i] == b[i] {
			idx++
		} else {
			break
		}
	}

	if idx < 0 {
		return ""
	}

	return string(a[:idx+1])
}

// 2nd solution (Binary Search)
// Runtime: 3ms
// Memory: 2.4MB
// Time Complexity: O(S*logN) -> S is the sum of all characters in strs slice, N is the number of characters in the shortest string
// Space Complexity: O(1)
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		minLen = min(minLen, len(strs[i]))
	}

	l, r := 0, minLen-1
	for l <= r {
		m := (l + r) / 2

		prefix := strs[0][:m+1]
		same := true
		for i := 1; i < len(strs); i++ {
			if !strings.HasPrefix(strs[i], prefix) {
				same = false
				break
			}
		}

		if !same {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	// r is the upper bound of the common prefix
	return strs[0][:r+1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
