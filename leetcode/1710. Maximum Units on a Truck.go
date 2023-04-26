package main

import "sort"

// Runtime: 27ms
// Memory Usage: 6.7 MB
// Time Complexity: O(nlogn) -> n is the length of boxTypes
// Space Complexity: O(1)
func maximumUnits(boxTypes [][]int, truckSize int) (ans int) {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	for i := 0; i < len(boxTypes); i++ {
		if boxTypes[i][0] > truckSize {
			ans += truckSize * boxTypes[i][1]
			return
		} else {
			ans += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		}
	}

	return
}
