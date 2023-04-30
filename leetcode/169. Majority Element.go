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
