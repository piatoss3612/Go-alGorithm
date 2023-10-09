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

	N       int
	corners [100001]int
	dp      [100001][3]int
)

// 난이도: Silver 1
// 메모리: 12380KB
// 시간: 36ms
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
		corners[i] = scanInt()
		dp[i][0] = -1
		dp[i][1] = -1
		dp[i][2] = -1
	}
}

func Solve() {
	fmt.Fprintln(writer, rec(1, 0))
}

func rec(here, successive int) int {
	// 기저 사례: N번째 시식 코너까지 지나온 경우
	if here == N+1 {
		return 0
	}

	ret := &dp[here][successive]
	if *ret != -1 {
		return *ret
	}

	*ret = 0

	// 이전에 시식 코너에 들리지 않았을 경우
	if successive == 0 {
		*ret = max(*ret, rec(here+1, 1)+corners[here])
	}

	// 이전에 시식 코너에 들렀을 경우
	if successive == 1 {
		*ret = max(*ret, rec(here+1, 2)+corners[here]/2)
	}

	// 이전에 시식 코너를 연속으로 두 번 들렀을 경우 또는 현재 시식 코너에 들어가지 않는 경우 (중복 케이스)
	*ret = max(*ret, rec(here+1, 0))

	return *ret
}

func max(a, b int) int {
	if a < b {
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
