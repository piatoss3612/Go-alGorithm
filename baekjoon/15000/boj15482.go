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
	S1, S2  string
)

// 15482번: 한글 LCS
// hhttps://www.acmicpc.net/problem/15482
// 난이도: 골드 5
// 메모리: 13872 KB
// 시간: 20 ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	S1 = scanString()
	S2 = scanString()
}

func Solve() {
	R1 := []rune(S1)
	R2 := []rune(S2)
	L1 := len(R1)
	L2 := len(R2)

	dp := make([][]int, L1+1)
	for i := range dp {
		dp[i] = make([]int, L2+1)
	}

	for i := 1; i <= L1; i++ {
		for j := 1; j <= L2; j++ {
			if R1[i-1] == R2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Fprintln(writer, dp[L1][L2])
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
