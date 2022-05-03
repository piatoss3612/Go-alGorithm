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
	dp      [1001][1001][2]int
)

// dp[x][y][0]: 길이가 x인 수열에서 인접한 비트가 y개이면서 0으로 끝나는 경우의 수
// dp[x][y][1]: 길이가 x인 수열에서 인접한 비트가 y개이면서 1로 끝나는 경우의 수
func init() {
	dp[1][0][0] = 1 // 0
	dp[1][0][1] = 1 // 1

	for i := 2; i <= 1000; i++ {

		// 인접한 비트가 0개인 경우와 i-1개인 경우
		dp[i][0][0] = dp[i-1][0][0] + dp[i-1][0][1]
		dp[i][0][1] = dp[i-1][0][0]
		dp[i][i-1][1] = 1 // 반드시 1이며 0으로 끝나는 경우는 존재하지 않는다

		for j := 1; j < i-1; j++ {
			dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
			dp[i][j][1] = dp[i-1][j-1][1] + dp[i-1][j][0]
		}

		/*
			점화식을 세우는 과정:

			예제 입력에 따라,

			dp[5][2][0]: 11100, 01110
			dp[5][2][1]: 00111, 10111, 11101, 11011

			dp[5][2][0]은 마지막 0을 제외하면 길이가 4인 수열에서 인접한 비트가 2개인 겨우의 수

			따라서, dp[5][2][0] = dp[4][2][0] + dp[4][2][1]이 된다.


			dp[5][2][1]은 다시 2가지 경우로 나뉘어진다.

			1. 마지막 1을 포함하고 길이가 4인 수열에서 인접한 비트가 1개이면서 1로 끝나는 경우의 수
			2. 마지막 1을 제외하고 길이가 4인 수열에서 인접한 비트가 2개이면서 0으로 끝나는 경우의 수

			따라서, dp[5][2][1] = dp[4][1][1] + dp[4][2][0]이 된다.

			점화식:
			dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
			dp[i][j][1] = dp[i-1][j-1][1] + dp[i-1][j][0]
		*/
	}
}

// 메모리: 12200KB
// 시간: 16ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	var a, b int

	for i := 1; i <= n; i++ {
		a, b = scanInt(), scanInt()
		fmt.Fprintln(writer, dp[a][b][0]+dp[a][b][1])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
