package main

// Runtime: 0ms
// Memory Usage: 2.1MB
// Time Complexity: O(n)
// Space Complexity: O(n)
func plusOne(digits []int) []int {
	n := len(digits) - 1

	digits[n] += 1

	for digits[n] > 9 {
		digits[n] -= 10

		if n == 0 {
			digits = append([]int{1}, digits...)
			break
		}

		n -= 1
		digits[n] += 1
	}

	return digits
}
