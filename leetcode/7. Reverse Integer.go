package main

// 7. Reverse Integer
// Runtime: 0ms
// Memory Usage: 2.1MB
// Time complexity: O(n) -> n is the number of digits
// Space complexity: O(1)
func reverse(x int) int {
	var negative bool

	// if x is negative, convert it to positive
	if x < 0 {
		negative = true
		x = -x
	}

	var result int

	// reverse the number
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}

	// if the number is negative, convert it to negative
	if negative {
		result = -result
	}

	// if the number is greater than 32-bit signed integer, return 0
	if result > 2147483647 || result < -2147483648 {
		return 0
	}

	return result
}