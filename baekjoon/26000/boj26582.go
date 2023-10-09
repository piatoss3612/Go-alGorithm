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

	T         int
	N, M      int
	artifacts []Artifact
	dp        []int
)

type Artifact struct {
	value  int
	weight int
}

// 난이도: Silver 1
// 메모리: 900KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
}

func Setup() {
	T = scanInt()

	for i := 1; i <= T; i++ {
		N, M = scanInt(), scanInt()
		artifacts = make([]Artifact, N+1)
		for j := 1; j <= N; j++ {
			artifacts[j] = Artifact{scanInt(), scanInt()}
		}
		dp = make([]int, M+1)
		Solve()
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		for j := M; j >= artifacts[i].weight; j-- {
			dp[j] = max(dp[j], dp[j-artifacts[i].weight]+artifacts[i].value) // dp[j] = max(현재 가치, j-artifacts[i].weight의 상태에서 artifacts[i]를 추가한 상태의 가치)
		}
	}

	fmt.Fprintln(writer, dp[M])
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
