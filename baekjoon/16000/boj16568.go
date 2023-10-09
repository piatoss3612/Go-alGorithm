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

	N, A, B int
	dp      [1000001]int
)

const INF = 987654321

// 난이도: Silver 1
// 메모리: 8736KB
// 시간: 16ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, A, B = scanInt(), scanInt(), scanInt()
}

func Solve() {
	// 메모리: 80716KB
	// 시간: 140ms
	// ans := rec(0)
	// fmt.Fprintln(writer, ans)

	for i := 1; i <= N; i++ {
		dp[i] = INF

		dp[i] = min(dp[i], dp[i-1]+1)
		if i-A-1 >= 0 {
			dp[i] = min(dp[i], dp[i-A-1]+1)
		}
		if i-B-1 >= 0 {
			dp[i] = min(dp[i], dp[i-B-1]+1)
		}
	}

	fmt.Fprintln(writer, dp[N])
}

func rec(here int) int {
	if here == N {
		return 0
	}

	ret := &dp[here]
	if *ret != 0 {
		return *ret
	}

	*ret = INF

	*ret = min(*ret, rec(here+1)+1)
	if here+A+1 <= N {
		*ret = min(*ret, rec(here+A+1)+1)
	}
	if here+B+1 <= N {
		*ret = min(*ret, rec(here+B+1)+1)
	}

	return *ret
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
