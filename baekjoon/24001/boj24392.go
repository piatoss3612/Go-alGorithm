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
	N, M    int
	bridge  [1002][1002]int
	dp      [1002][1002]int
)

const MOD = 1000000007

// 난이도: Silver 1
// 메모리: 16572KB
// 시간: 68ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			bridge[i][j] = scanInt()
		}
	}
}

func Solve() {
	for i := 1; i <= M; i++ {
		dp[N][i] = bridge[N][i]
	}

	for i := N - 1; i >= 1; i-- {
		for j := 1; j <= M; j++ {
			if bridge[i][j] == 1 {
				dp[i][j] = (dp[i+1][j-1] + dp[i+1][j] + dp[i+1][j+1]) % MOD
			}
		}
	}

	ans := 0
	for i := 1; i <= M; i++ {
		ans = (ans + dp[1][i]) % MOD
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
