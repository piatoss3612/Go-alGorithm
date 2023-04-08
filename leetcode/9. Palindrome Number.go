package main

import "strconv"

// 1st solution
// Runtime: 28ms
// Memory: 6.4MB
// Time complexity: O(log10(n))
// Space complexity: O(log10(n))? O(1)? (not sure)
func isPalindrome(x int) bool {
	// -121 != 121-
	if x < 0 {
		return false
	}

	// get digits of x
	digits := []int{}
	for x > 0 {
		r := x % 10
		x /= 10
		digits = append(digits, r)
	}

	// digits has reversed order of x though it doesn't matter
	// check if x is palindrome
	for i, j := 0, len(digits)-1; i <= j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

// 2nd solution
// Runtime: 15ms
// Memory: 4.8MB
// Time complexity: O(log10(n)) -> O(log10(n)/2)
// Space complexity: O(log10(n))? O(1)? (not sure)
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	digits := strconv.Itoa(x)

	for i, j := 0, len(digits)-1; i <= j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}

	return true
}

// 3rd solution
// Runtime: 20ms
// Memory: 4.6MB
// Time complexity: O(log10(n))
// Space complexity: O(1)
func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}

	y := x

	reversed := 0
	for y > 0 {
		reversed = reversed*10 + (y % 10)
		y /= 10
	}

	return reversed == x
}

// 2nd solution is the fastest one
// 3rd solution is the most memory efficient one
