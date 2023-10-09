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
	prey [21]int
	dp   [21]int
)

// 난이도: Silver1
// 메모리: 908KB
// 시간: 4ms
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
		prey[i] = scanInt()
		dp[i] = -1
	}
}

func Solve() {
	fmt.Fprintln(writer, rec(1))
}

func rec(n int) int {
	// 기저 사례: 더 이상 먹이가 없는 경우
	if n == N+1 {
		return 0
	}

	ret := &dp[n]
	if *ret != -1 {
		return *ret
	}

	*ret = 0
	*ret = max(*ret, rec(n+1)) // n번째 먹이를 먹지 않는 경우

	sum := 0
	for i := n; i <= N; i++ {
		sum += prey[i]
		if sum >= M {
			*ret = max(*ret, rec(i+1)+sum-M) // n번째 먹이부터 i번째 먹이까지 먹는 경우: 탈피 에너지가 sum-M만큼 증가
			break
		}
	}

	return *ret
}

func max(a, b int) int {
	if a > b {
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
