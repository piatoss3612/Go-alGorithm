package main

// goal: check if a number is a perfect square
// constraints: not use any built-in library function such as sqrt

// 1st solution - binary search
// Runtime: 1ms
// Memory Usage: 1.9MB
// Time Complexity: O(logn)
// Space Complexity: O(1)
func isPerfectSquare(num int) bool {
	l, r := 1, num
	for l <= r {
		m := (l + r) / 2
		if m*m <= num {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return r*r == num
}

// 2nd solution - brute force
// Runtime: 1ms
// Memory Usage: 1.9MB
// Time Complexity: O(sqrt(n))
// Space Complexity: O(1)
func isPerfectSquare2(num int) bool {
	n := 1
	for i := 2; i*i <= num; i++ {
		n = i
	}

	return n*n == num
}
