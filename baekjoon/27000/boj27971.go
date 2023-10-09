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

	N, M, A, B  int
	constraints [][2]int
	dp          [100001]int
)

const INF = 987654321

// 난이도: Silver 1
// 메모리: 1696KB
// 시간: 20ms
// 분류: 다이나믹 프로그래밍, 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, A, B = scanInt(), scanInt(), scanInt(), scanInt()
	constraints = make([][2]int, M)
	for i := 0; i < M; i++ {
		constraints[i] = [2]int{scanInt(), scanInt()}
	}

	for i := 0; i <= N; i++ {
		dp[i] = INF
	}
}

func Solve() {
	dp[0] = 0
	for i := 1; i <= N; i++ {
		if i-A >= 0 && dp[i-A] != INF {
			dp[i] = dp[i-A] + 1
		}

		if i-B >= 0 && dp[i-B] != INF {
			dp[i] = min(dp[i], dp[i-B]+1)
		}

		for _, constraint := range constraints {
			if i >= constraint[0] && i <= constraint[1] {
				dp[i] = INF
				break
			}
		}
	}

	if dp[N] == INF {
		fmt.Fprintln(writer, -1)
		return
	}
	fmt.Fprintln(writer, dp[N])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
