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
	D       int
	dp      [2][8]int
)

const MOD = 1000000007

// 난이도: Silver 1
// 메모리: 916KB
// 시간: 8ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	D = scanInt()
}

func Solve() {
	dp[0][0] = 1

	for i := 1; i <= D; i++ {
		dp[1][0] = dp[0][1] + dp[0][2]
		dp[1][1] = dp[0][0] + dp[0][2] + dp[0][3]
		dp[1][2] = dp[0][0] + dp[0][1] + dp[0][3] + dp[0][4]
		dp[1][3] = dp[0][1] + dp[0][2] + dp[0][4] + dp[0][5]
		dp[1][4] = dp[0][2] + dp[0][3] + dp[0][5] + dp[0][6]
		dp[1][5] = dp[0][3] + dp[0][4] + dp[0][7]
		dp[1][6] = dp[0][4] + dp[0][7]
		dp[1][7] = dp[0][5] + dp[0][6]

		for i := 0; i <= 7; i++ {
			dp[0][i] = dp[1][i] % MOD
		}
	}

	fmt.Fprintln(writer, dp[0][0])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
