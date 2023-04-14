package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution: Recursive
// Runtime: 0ms
// Memory: 2.1MB
// Time Complexity: O(n) -> n is the number of nodes in the tree
// Space Complexity: O(n)
func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	result = append(result, inorderTraversal(root.Left)...)

	result = append(result, root.Val)

	result = append(result, inorderTraversal(root.Right)...)

	return result
}
