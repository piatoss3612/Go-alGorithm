package level02

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/12945
// 분류: 다이나믹 프로그래밍
func solution(n int) int {
	dp := make([]int, n+1)

	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1234567
	}

	return dp[n]
}
