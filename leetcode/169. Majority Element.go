package main

// 1st solution - map
// Runtime: 18ms
// Memory Usage: 6.3 MB
// Time complexity: O(n)
// Space complexity: O(n)
func majorityElement(nums []int) int {
	cnts := make(map[int]int)

	for _, num := range nums {
		cnts[num] += 1
		if cnts[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}

// 2nd solution - Boyer-Moore Voting Algorithm
// Runtime: 17ms
// Memory Usage: 6.2 MB
// Time complexity: O(n)
// Space complexity: O(1)
// Ref: https://leetcode.com/problems/majority-element/solutions/3220196/beats-99-87-with-go/
func majorityElement2(nums []int) int {
	candidate := nums[0]
	cnt := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			cnt += 1
		} else {
			cnt -= 1
		}

		if cnt == 0 {
			candidate = nums[i]
			cnt = 1
		}
	}

	return candidate
}
