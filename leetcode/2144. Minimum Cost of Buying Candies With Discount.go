package main

import (
	"sort"
)

// 1st solution
// Runtime: 6ms
// Memory Usage: 2.7MB
// Time Complexity: O(nlogn) -> sort
// Space Complexity: O(1)
func minimumCost(cost []int) (sum int) {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})

	for len(cost) >= 3 {
		sum += cost[0] + cost[1]
		cost = cost[3:]
	}

	for _, c := range cost {
		sum += c
	}

	return
}

// 2nd solution
// Runtime: 3ms
// Runtime: 2.7MB
// Time Complexity: O(nlogn) -> sort
// Space Complexity: O(1)
func minimumCost2(cost []int) (sum int) {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})

	for i, c := range cost {
		if i%3 == 2 {
			continue
		}
		sum += c
	}

	return
}
