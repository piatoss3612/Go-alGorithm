package main

import "fmt"

// 1277. Count Square Submatrices with All Ones
// Runtime: 132ms
// Memory Usage: 7.4MB
// Time complexity: O(m*n*min(m,n))
// Space complexity: O(m*n)
// Category: Partial Sum
func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		sum[i] = make([]int, n+1)
	}

	// sum[i][j] = sum of matrix[0][0] to matrix[i-1][j-1]
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = matrix[i-1][j-1] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}

	cnt := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// k is the length of the square
			for k := 1; k <= min(m-i+1, n-j+1); k++ {
				if sum[i+k-1][j+k-1]-sum[i-1][j+k-1]-sum[i+k-1][j-1]+sum[i-1][j-1] == k*k {
					cnt++
				}
			}
		}
	}

	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	mat := [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}
	fmt.Println(countSquares(mat))
}
