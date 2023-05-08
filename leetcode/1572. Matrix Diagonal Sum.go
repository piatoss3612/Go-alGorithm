package main

// 1st solution
// Runtime: 20ms
// Memory Usage: 5.7MB
// Time complexity: O(n*n)
// Space complexity: O(1)
func diagonalSum(mat [][]int) (sum int) {
	n := len(mat)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				sum += mat[i][j]
				mat[i][j] = 0
			}
		}
	}

	return
}

// 2nd solution
// Runtime: 14ms
// Memory Usage: 5.6MB
// Time complexity: O(n)
// Space complexity: O(1)
func diagonalSum2(mat [][]int) (sum int) {
	n := len(mat)

	for i := 0; i < n; i++ {
		sum += mat[i][i]
		if i != n-i-1 {
			sum += mat[i][n-i-1]
		}
	}
	return
}
