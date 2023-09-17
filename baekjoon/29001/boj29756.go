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

	N, K int
	s    [1001]int
	h    [1001]int
	dp   [1001][101]int
)

// 난이도: Gold 4
// 메모리: 1856KB
// 시간: 8ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		s[i] = scanInt()
	}
	for i := 1; i <= N; i++ {
		h[i] = scanInt()
	}

	for i := 0; i <= 1000; i++ {
		for j := 0; j <= 100; j++ {
			dp[i][j] = -1
		}
	}
}

func Solve() {
	ans := rec(0, 100) // 체력 100으로 시작
	fmt.Fprintln(writer, ans)
}

func rec(here, hp int) int {
	// 기저 사례: 모든 구간을 지나왔을 때
	if here == N+1 {
		return 0
	}

	ret := &dp[here][hp]
	if *ret != -1 {
		return *ret
	}

	hp += K // K만큼 회복
	if hp > 100 { // 최대 체력은 100
		hp = 100
	}

	// here번째 구간을 건너뛰는 경우
	*ret = rec(here+1, hp)

	// here번째 구간을 플레이하는 경우: 체력이 0미만이 되지 않는 경우에만 가능
	if hp-h[here] >= 0 {
		*ret = max(*ret, rec(here+1, hp-h[here])+s[here])
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