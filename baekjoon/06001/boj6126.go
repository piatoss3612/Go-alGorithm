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

	V, N  int
	coins [26]int
	dp    [10001]int
)

// 난이도: Silver 1
// 메모리: 1068KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	V, N = scanInt(), scanInt()
	for i := 1; i <= V; i++ {
		coins[i] = scanInt()
	}
}

func Solve() {
	dp[0] = 1
	for i := 1; i <= V; i++ {
		for j := coins[i]; j <= N; j++ {
			// dp[j+coins[i]] += dp[j]를 하면 중복이 발생
			dp[j] += dp[j-coins[i]] // dp[j]는 j원을 만들 수 있는 경우의 수
		}
	}
	fmt.Fprintln(writer, dp[N])
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
