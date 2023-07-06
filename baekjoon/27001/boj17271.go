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

	N, M int
	dp   [10001][10001]int // 조합의 경우의 수를 저장할 배열
)

const MOD = 1000000007

// 난이도: Silver 1
// 메모리: 430600KB
// 시간: 356ms
// 분류: 다이나믹 프로그래밍, 조합론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
}

func Solve() {
	for i := 1; i <= N; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				dp[i][j] = 1
			} else {
				dp[i][j] = (dp[i-1][j-1] + dp[i-1][j]) % MOD
			}
		}
	}

	total := 0

	for i := 0; i*M <= N; i++ {
		total = (total + dp[N-i*M+i][i]) % MOD // dp[N-i*M+i][i]: M을 i개 사용하고 나머지 N-i*M개를 1~M으로 만드는 경우의 수
	}

	fmt.Fprintln(writer, total)
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
