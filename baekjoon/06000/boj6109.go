package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, C    int
	coins   []int
)

// 6109번: Dairy Queen
// https://www.acmicpc.net/problem/6109
// 난이도: 골드 5
// 메모리: 900 KB
// 시간: 4 ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, C = scanInt(), scanInt()
	coins = make([]int, C+1)
	for i := 1; i <= C; i++ {
		coins[i] = scanInt()
	}
}

func Solve() {
	dp := make([][]int, C+1)
	for i := 0; i <= C; i++ {
		dp[i] = make([]int, N+1)
	}

	dp[0][0] = 1
	for i := 1; i <= C; i++ {
		for j := 0; j <= N; j++ {
			dp[i][j] = dp[i-1][j]
			if j-coins[i] >= 0 {
				dp[i][j] += dp[i][j-coins[i]]
			}
		}
	}

	fmt.Fprintln(writer, dp[C][N])
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
