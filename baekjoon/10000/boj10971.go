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
	cost    [11][11]int
	visited [11]bool

	ans = 987654321
)

// 난이도: Silver 2
// 메모리: 908KB
// 시간: 24ms
// 분류: 브루트포스 알고리즘, 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			cost[i][j] = scanInt()
		}
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		rec(N, i, i, 0)
	}

	fmt.Fprintln(writer, ans)
}

func rec(remain, start, here, total int) {
	// 모든 도시를 방문했고, 시작점으로 돌아온 경우
	if remain == 0 && start == here {
		ans = min(ans, total)
		return
	}

	// 남은 도시를 모두 방문해도 ans보다 큰 경우
	if total > ans {
		return
	}

	// 현재 도시에서 갈 수 있는 도시를 차례대로 방문
	for i := 1; i <= N; i++ {
		if visited[i] || cost[here][i] == 0 {
			continue
		}

		visited[i] = true
		rec(remain-1, start, i, total+cost[here][i])
		visited[i] = false
	}
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
