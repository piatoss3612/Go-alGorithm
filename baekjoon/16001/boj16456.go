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

	N  int
	dp [50001]int
)

const MOD = 1000000009

// 난이도: Silver 1
// 메모리: 1308KB
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
}

func Solve() {
	// 첫번째 섬에서 시작한다고 가정
	dp[1] = 1 // 섬이 1개일 때는 1가지 => ()
	dp[2] = 1 // 섬이 2개일 때는 1가지 => (1)
	dp[3] = 2 // 섬이 3개일 때는 2가지 => (1, 1), (2, -1)
	// dp[4] = 3 // 섬이 4개일 때는 3가지 => (1, 1, 1), (1, 2, -1), (2, 1, -1)
	// dp[5] = 4 // 섬이 5개일 때는 4가지 => (1, 1, 1, 1), (1, 1, 2, -1), (1, 2, 1, -1), (2, -1, 2, 1)

	// 점화식: dp[i] = dp[i-1] + dp[i-3]

	for i := 4; i <= N; i++ {
		dp[i] = (dp[i-1] + dp[i-3]) % MOD
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
