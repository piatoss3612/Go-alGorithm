package level02

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/12914
// 분류: 다이나믹 프로그래밍
func solution(n int) int64 {
	dp := make([]int64, n+1)

	dp[1] = 1

	if n > 1 {
		dp[2] = 2
	}

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1234567
	}

	return dp[n]
}
