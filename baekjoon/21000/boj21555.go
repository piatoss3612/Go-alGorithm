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

	N, K int
	A    [200001]int
	B    [200001]int
	dp   [200001][2]int
)

// 난이도: Silver 2
// 메모리: 12220KB
// 시간: 112ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		A[i] = scanInt()
	}
	for i := 1; i <= N; i++ {
		B[i] = scanInt()
	}
}

func Solve() {
	dp[1][0], dp[1][1] = A[1], B[1]
	for i := 2; i <= N; i++ {
		dp[i][0] = min(dp[i-1][0], dp[i-1][1]+K) + A[i]
		dp[i][1] = min(dp[i-1][1], dp[i-1][0]+K) + B[i]
	}

	fmt.Fprintln(writer, min(dp[N][0], dp[N][1]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
