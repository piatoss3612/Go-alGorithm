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

	N   int
	arr [1001]int
	dp  [1001]int
)

// 난이도: Silver 1
// 메모리: 924KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		dp[i] = arr[i]
		for j := 1; j < i; j++ {
			if arr[i] < arr[j] {
				dp[i] = max(dp[i], dp[j]+arr[i])
			}
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		ans = max(ans, dp[i])
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
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
