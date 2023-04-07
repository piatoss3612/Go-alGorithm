package main

import "sort"

// Runtime: 7ms
// Memory: 5.7MB
// Time complexity: O(n) -> hash map
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	has := make(map[int]int) // key: num, value: index, used to check if j is in nums
	for i, num := range nums {
		has[num] = i + 1 // i+1 is used to avoid 0
	}

	result := make([]int, 0, 2)

	for i, num := range nums {
		j := target - num
		_, ok := has[j]
		// j is in nums and j is not num
		if ok && has[j]-1 != i {
			result = append(result, i, has[j]-1)
			break
		}
	}

	sort.Ints(result) // sort result
	return result
}

// Runtime: 35ms
// Memory: 3.7MB
// Time complexity: O(n^2) -> brute force
// Space complexity: O(1)
func twoSumBruteForce(nums []int, target int) []int {
	result := make([]int, 0, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}
