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
	N, T    int
	dp      [1001]int
)

// 29704번: 벼락치기
// https://www.acmicpc.net/problem/29704
// 난이도: 골드 5
// 메모리: 868 KB
// 시간: 8 ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, T = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		dp[i] = -1
	}
}

func Solve() {
	total := 0
	for i := 1; i <= N; i++ {
		d, m := scanInt(), scanInt()
		total += m

		for j := T; j >= d; j-- {
			if dp[j-d] != -1 {
				dp[j] = max(dp[j], dp[j-d]+m)
			}
		}
	}

	ans := total
	for i := 1; i <= T; i++ {
		ans = min(ans, total-dp[i])
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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