package main

// Runtime: 1ms
// Memory Usage: 1.9MB
// Time complexity: O(1)
// Space complexity: O(1)
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return min(rec1[2], rec2[2]) > max(rec1[0], rec2[0]) && min(rec1[3], rec2[3]) > max(rec1[1], rec2[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
