package main

// 1st solution - dynamic programming
// Runtime: 0ms
// Memory: 2MB
// Time complexity: O(n)
// Space complexity: O(n)
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 2nd solution - recursive with memoization
// Runtime: 0ms
// Memory: 2MB
// Time complexity: O(n)
// Space complexity: O(n)
var dp = make([]int, 46)

func climbStairs2(n int) int {
	if n <= 1 {
		return 1
	}

	if dp[n] != 0 {
		return dp[n]
	}

	dp[n] = climbStairs(n-1) + climbStairs(n-2)
	return dp[n]
}

// 3rd solution - dynamic programming with O(1) space
// Runtime: 1ms
// Memory: 1.9MB
// Time complexity: O(n)
// Space complexity: O(1)
func climbStairs3(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
