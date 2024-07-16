package main

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/43165
// 분류: 다이나믹 프로그래밍
func solution(numbers []int, target int) int {
	dp := make([][2001]int, len(numbers)+1)
	dp[0][1000] = 1

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < 2001; j++ {
			if dp[i][j] == 0 {
				continue
			}

			dp[i+1][j+numbers[i]] += dp[i][j]
			dp[i+1][j-numbers[i]] += dp[i][j]
		}
	}

	return dp[len(numbers)][target+1000]
}

func main() {
	println(solution([]int{1, 1, 1, 1, 1}, 3)) // 5
}
