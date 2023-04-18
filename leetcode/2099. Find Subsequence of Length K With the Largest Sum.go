package main

import (
	"sort"
)

// Runtime: 3ms
// Memory Usage: 3.4MB
// Time Complexity: O(nlogn) -> sort
// Space Complexity: O(n)
func maxSubsequence(nums []int, k int) []int {
	type IndexedNum struct {
		index int
		num   int
	}

	inums := make([]IndexedNum, 0, len(nums))

	for i, n := range nums {
		inums = append(inums, IndexedNum{
			index: i,
			num:   n,
		})
	}

	sort.Slice(inums, func(i, j int) bool {
		return inums[i].num < inums[j].num
	})

	inums = inums[len(inums)-k:]

	sort.Slice(inums, func(i, j int) bool {
		return inums[i].index < inums[j].index
	})

	ans := make([]int, 0, k)

	for _, in := range inums {
		ans = append(ans, in.num)
	}

	return ans
}
