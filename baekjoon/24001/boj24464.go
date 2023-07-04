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
	dp [200001][5]int
)

const MOD = 1000000007

// 난이도: Silver 1
// 메모리: 8672KB
// 시간: 12ms
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
	dp[1][0] = 1
	dp[1][1] = 1
	dp[1][2] = 1
	dp[1][3] = 1
	dp[1][4] = 1
	for i := 2; i <= N; i++ {
		dp[i][0] += (dp[i-1][1] + dp[i-1][2] + dp[i-1][3] + dp[i-1][4]) % MOD
		dp[i][1] += (dp[i-1][0] + dp[i-1][3] + dp[i-1][4]) % MOD
		dp[i][2] += (dp[i-1][0] + dp[i-1][4]) % MOD
		dp[i][3] += (dp[i-1][0] + dp[i-1][1]) % MOD
		dp[i][4] += (dp[i-1][0] + dp[i-1][1] + dp[i-1][2]) % MOD
	}

	fmt.Fprintln(writer, (dp[N][0]+dp[N][1]+dp[N][2]+dp[N][3]+dp[N][4])%MOD)
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
