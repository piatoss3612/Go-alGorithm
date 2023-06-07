package main

import "sort"

// Runtime: 5ms
// Memory Usage: 2.5MB
// Time complexity: O(nlogn)
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	diff := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true
}
