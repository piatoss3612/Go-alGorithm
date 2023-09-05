package main

// 1461. Check If a String Contains All Binary Codes of Size K
// Runtime: 134ms
// Memory Usage: 24.94MB
// Time Complexity: O(n)
// Space Complexity: O(2^k)
func hasAllCodes(s string, k int) bool {
	check := make(map[string]bool)

	for i := 0; i <= len(s)-k; i++ {
		check[s[i:i+k]] = true
	}

	return len(check) == 1<<k
}
