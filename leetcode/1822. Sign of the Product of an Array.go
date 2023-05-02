package main

// Runtime: 3ms
// Memory Usage: 3.3 MB
// Time complexity: O(n)
// Space complexity: O(1)
func arraySign(nums []int) int {
	product := 1

	for i := 0; i < len(nums); i++ {
		product *= signFunc(nums[i])
	}

	return signFunc(product)
}

func signFunc(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}
