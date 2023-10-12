package main

// 2439. Minimize Maximum of Array
// Runtime: 95 ms
// Memory: 8.64 MB
// Time Complexity: O(n log n)
// Topic: Binary Search
func minimizeArrayValue(nums []int) int {
	r := 0
	for _, v := range nums {
		if v > r {
			r = v
		}
	}

	l := 0

	for l <= r {
		mid := (l + r) / 2

		if check(nums, mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func check(nums []int, mid int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum > mid*(i+1) {
			return false
		}
	}

	return true
}