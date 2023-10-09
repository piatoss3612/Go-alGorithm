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
	arr  [501][501]int
	dp   [501][501]int
)

// 난이도: Gold 5
// 메모리: 6176KB
// 시간: 376ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

const INF = 987654321

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			arr[i][j] = scanInt()
			dp[i][j] = INF
		}
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			for k := 1; k <= M; k++ {
				// 같은 열은 건너뛰기
				if j == k {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+arr[i][j])
			}
		}
	}

	ans := INF
	for i := 1; i <= M; i++ {
		ans = min(ans, dp[N][i])
	}

	fmt.Fprintln(writer, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
