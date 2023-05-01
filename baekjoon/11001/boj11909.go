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

	N     int
	board [2223][2223]int
	dp    [2223][2223]int
)

const INF = 987654321

// 난이도: Gold 5
// 메모리: 83796KB
// 시간: 484ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		dp[i][0] = INF
		for j := 1; j <= N; j++ {
			board[i][j] = scanInt()
			dp[i][j] = INF
		}
	}
	dp[0] = dp[1]
}

func Solve() {
	dp[1][1] = 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if i+1 <= N {
				if board[i+1][j] >= board[i][j] {
					dp[i+1][j] = min(dp[i+1][j], dp[i][j]+board[i+1][j]-board[i][j]+1)
				} else {
					dp[i+1][j] = min(dp[i+1][j], dp[i][j])
				}
			}

			if j+1 <= N {
				if board[i][j+1] >= board[i][j] {
					dp[i][j+1] = min(dp[i][j+1], dp[i][j]+board[i][j+1]-board[i][j]+1)
				} else {
					dp[i][j+1] = min(dp[i][j+1], dp[i][j])
				}
			}
		}
	}

	fmt.Fprintln(writer, dp[N][N])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
