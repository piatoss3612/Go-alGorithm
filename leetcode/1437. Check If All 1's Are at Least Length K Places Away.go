package main

// 1st solution
// Runtime: 73ms
// Memory Usage: 7.8MB
// Time complexity: O(n)
// Space complexity: O(1)
func kLengthApart(nums []int, k int) bool {
	n := len(nums)
	cnt := 0
	prev := -1

	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			prev = i
			cnt += 1
			break
		}
	}

	dist := 0

	for i := prev + 1; i < n; i++ {
		if nums[i] == 0 {
			dist += 1
			continue
		}

		if dist < k {
			return false
		}

		prev = i
		dist = 0
		cnt += 1
	}

	return true
}

// 2nd solution
// Runtime: 61ms
// Memory Usage: 7.6MB
// Time complexity: O(n)
// Space complexity: O(1)
func kLengthApart2(nums []int, k int) bool {
	prev := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if prev != -1 && i-prev-1 < k {
				return false
			}
			prev = i
		}
	}

	return true
}
