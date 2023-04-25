package main

const INF = 987654321

// Runtime: 15ms
// Memory Usage: 6.6 MB
// Time complexity: O(n*m)
// Space complexity: O(n+m)
func luckyNumbers(matrix [][]int) (res []int) {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		row[i] = INF
		for j := 0; j < len(matrix[0]); j++ {
			row[i] = min(row[i], matrix[i][j])
			col[j] = max(col[j], matrix[i][j])
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == row[i] && col[j] == row[i] {
				res = append(res, row[i])
				break
			}
		}
	}
	return
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
