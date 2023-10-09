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

	N, M      int
	bracelets [][2]int
	dp        []int
)

// 난이도: Gold 5
// 메모리: 1212KB
// 시간: 100ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	bracelets = make([][2]int, N+1)
	dp = make([]int, M+1)

	for i := 1; i <= N; i++ {
		bracelets[i][0], bracelets[i][1] = scanInt(), scanInt() // 무게, 가치
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		// i번째 팔찌를 담았을 때 무게가 j인 경우의 최대 가치를 구한다.
		for j := M; j >= bracelets[i][0]; j-- {
			dp[j] = max(dp[j], dp[j-bracelets[i][0]]+bracelets[i][1])
		}
	}
	fmt.Fprintln(writer, dp[M]) // 최대 가치를 출력한다.
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
