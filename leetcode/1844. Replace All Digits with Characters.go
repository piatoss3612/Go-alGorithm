package main

// Runtime: 0ms
// Memory Usage: 2MB
// Time complexity: O(n) -> n is the length of s
// Space complexity: O(n)
func replaceDigits(s string) string {
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if i%2 == 1 {
			b[i] = b[i-1] + (b[i] - '0')
		}
	}

	return string(b)
}
