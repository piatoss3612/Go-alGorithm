package main

// 2374. Node With Highest Edge Score
// Runtime: 143ms
// Memory Usage: 8.38MB
// Time: O(n) n is the number of nodes
// Space: O(n) n is the number of nodes
func edgeScore(edges []int) int {
	arr := make([]int, len(edges))

	for i := 0; i < len(edges); i++ {
		arr[edges[i]] += i
	}

	maxNode := 0

	for i := 1; i < len(arr); i++ {
		if arr[maxNode] < arr[i] {
			maxNode = i
		}
	}

	return maxNode
}
