package main

import "sort"

// solution - Brute force and sort
// Runtime: 13ms
// Memory Usage: 7 MB
// Time complexity: O(nlogn) -> n is the length of arr
// Space complexity: O(10001) -> digits array
func sortByBits(arr []int) []int {
	digits := make([]int, 10001)
	digits[1] = 1
	for i := 2; i <= 10000; i++ {
		if i%2 == 0 {
			digits[i] = digits[i/2]
		} else {
			digits[i] = digits[i-1] + digits[1]
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		if digits[arr[i]] == digits[arr[j]] {
			return arr[i] < arr[j]
		}
		return digits[arr[i]] < digits[arr[j]]
	})

	return arr
}
