package main

// 1st solution - brute force
// Runtime: 16ms
// Memory: 2.1MB
// Time complexity: O(√n)
// Space complexity: O(1)
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	result := 1
	for i := 2; i*i <= x; i++ {
		result = i
	}
	return result
}

// 2nd solution - binary search
// Runtime: 0ms
// Memory: 2.1MB
// Time complexity: O(logn)
// Space complexity: O(1)
func mySqrt2(x int) int {
	l, r := 0, x
	for l <= r {
		m := (l + r) / 2

		square := m * m
		if square == x {
			return m
		}

		if square < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return r
}
