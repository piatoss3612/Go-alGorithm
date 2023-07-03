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

	N, M  int
	items [301][301]int
	dp    [301][301]int
)

const INF = 987654321

// 난이도: Silver 1
// 메모리: 2392KB
// 시간: 16ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			items[i][j] = scanInt()
			dp[i][j] = INF
		}
	}
}

func Solve() {
	dp[1][1] = 0 // (1, 1)에서 시작
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			// (i, j)에 도달할 수 없는 경우
			if dp[i][j] == INF {
				continue
			}

			// 부스터를 1개에서 items[i][j]개 사용할 수 있음
			for k := 1; k <= items[i][j]; k++ {
				// (i+k, j) 또는 (i, j+k)로 이동하는 경우
				if i+k <= N {
					dp[i+k][j] = min(dp[i+k][j], dp[i][j]+1)
				}
				if j+k <= M {
					dp[i][j+k] = min(dp[i][j+k], dp[i][j]+1)
				}
			}
		}
	}

	fmt.Fprintln(writer, dp[N][M])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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