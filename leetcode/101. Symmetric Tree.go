package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1st solution - Recursive
// Runtime: 0ms
// Memory Usage: 2.9MB
// Time complexity: O(n) -> n is the number of nodes
// Space complexity: O(d) -> d is the depth of the tree
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSame(root.Left, root.Right)
}

func isSame(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	return left.Val == right.Val &&
		isSame(left.Left, right.Right) &&
		isSame(left.Right, right.Left)
}

// 2nd solution - Iterative
// Runtime: 0ms
// Memory Usage: 3MB
// Time complexity: O(n) -> n is the number of nodes
// Space complexity: O(n)
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := []*TreeNode{}
	q = append(q, root.Left, root.Right)

	for len(q) >= 2 {
		left, right := q[0], q[1]
		q = q[2:]

		if left == nil || right == nil {
			if left != right {
				return false
			} else {
				continue
			}
		}

		if left.Val != right.Val {
			return false
		}

		q = append(q, left.Left, right.Right, left.Right, right.Left)
	}

	return true
}
