package main

// 240. Search a 2D Matrix II
// Runtime: 21ms
// Memory Usage: 6.82MB
// Time Complexity: O(n + m) -> n: number of rows, m: number of columns
// Space Complexity: O(1)
// Reference: https://cs.stackexchange.com/questions/98575/proving-correctness-of-search-algorithms
func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1 // start from top right corner

	for i < n && j >= 0 {
		// if matrix[i][j] is equal to target, then we found target
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			// if matrix[i][j] is greater than target, then all elements in jth column are greater than target
			// therefore, we can skip jth column
			j--
		} else {
			// if matrix[i][j] is less than target, then all elements in ith row (0 ~ jth column) are less than target
			// therefore, we can skip ith row
			i++
		}
	}

	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7, 9}, {2, 4, 6, 8, 10}, {11, 13, 15, 17, 19}, {12, 14, 16, 18, 20}, {21, 22, 23, 24, 25}}
	target := 13

	println(searchMatrix(matrix, target))
}
