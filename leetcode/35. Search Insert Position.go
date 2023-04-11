package main

// Runtime: 2ms
// Memory: 3MB
// Time Complexity: O(logn)
// Space Complexity: O(1)
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		if target <= nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return l // lower bound: min position to insert target
}
