package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1st solution - recursive
// Runtime: 3ms
// Memory Usage: 4.2MB
// Time Complexity: O(n) - n is the number of nodes
// Space Complexity: O(1)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 1
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		ans += left
	} else {
		ans += right
	}

	return ans
}

// 2nd solution - recursive, cleaner but slower
// Runtime: 6ms
// Memory Usage: 4.3MB
// Time Complexity: O(n) - n is the number of nodes
// Space Complexity: O(1)
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
