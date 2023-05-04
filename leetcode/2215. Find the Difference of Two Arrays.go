package main

// Runtime: 27ms
// Memory Usage: 6.8MB
// Time Complexity: O(n)
// Space Complexity: O(n)
func findDifference(nums1 []int, nums2 []int) (result [][]int) {
	result = append(result, getDifference(nums1, nums2), getDifference(nums2, nums1))
	return
}

func getDifference(a, b []int) []int {
	result := make([]int, 0, len(a))
	has := make([]bool, 2001)

	for i := 0; i < len(b); i++ {
		has[b[i]+1000] = true
	}

	for i := 0; i < len(a); i++ {
		if !has[a[i]+1000] {
			result = append(result, a[i])
			has[a[i]+1000] = true
		}
	}

	return result
}
