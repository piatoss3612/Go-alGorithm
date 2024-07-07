package main

import "fmt"

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/77485
// 분류: 구현, 그래프
func solution(rows int, columns int, queries [][]int) []int {
	matrix := make([][]int, rows+1)
	for i := 1; i <= rows; i++ {
		matrix[i] = make([]int, columns+1)
		for j := 1; j <= columns; j++ {
			matrix[i][j] = (i-1)*columns + j
		}
	}

	answer := make([]int, 0, len(queries))

	for _, query := range queries {
		x1, y1, x2, y2 := query[0], query[1], query[2], query[3]

		temp := matrix[x1][y2]
		minVal := temp

		// 시계방향 회전

		// 1. 위쪽
		for i := y2; i > y1; i-- {
			matrix[x1][i] = matrix[x1][i-1]
			minVal = min(minVal, matrix[x1][i])
		}

		// 2. 왼쪽
		for i := x1; i < x2; i++ {
			matrix[i][y1] = matrix[i+1][y1]
			minVal = min(minVal, matrix[i][y1])
		}

		// 3. 아래쪽
		for i := y1; i < y2; i++ {
			matrix[x2][i] = matrix[x2][i+1]
			minVal = min(minVal, matrix[x2][i])
		}

		// 4. 오른쪽
		for i := x2; i > x1; i-- {
			matrix[i][y2] = matrix[i-1][y2]
			minVal = min(minVal, matrix[i][y2])
		}

		matrix[x1+1][y2] = temp

		answer = append(answer, minVal)
	}

	return answer
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	answer := solution(6, 6, [][]int{{2, 2, 5, 4}, {3, 3, 6, 6}, {5, 1, 6, 3}})
	fmt.Println(answer) // [8, 10, 25]
	answer = solution(3, 3, [][]int{{1, 1, 2, 2}, {1, 2, 2, 3}, {2, 1, 3, 2}, {2, 2, 3, 3}})
	fmt.Println(answer) // [1, 1, 5, 3]
}
