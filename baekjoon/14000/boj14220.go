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
	path    [501][501]int
	dp      [501][501]int
)

const INF = 987654321

// 난이도: Gold 3
// 메모리: 6772KB
// 시간: 456ms
// 분류: 다이나믹 프로그래밍, 깊이 우선 탐색
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
			path[i][j] = scanInt()
		}
		path[0][i] = 1 // 가상의 0번 도시와 i번 도시가 이동거리 1로 연결되어 있다고 가정
	}
}

func Solve() {
	ans := rec(0, 0)
	// 최소 경로가 존재하지 않는 경우
	if ans == INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans-1) // 최소 경로가 존재할 경우 처음에 0번 도시와 연결된 비용 1을 제거
	}
}

func rec(turn, here int) int {
	// 기저 사례: 최단 거리를 찾은 경우
	if turn == N {
		return 0
	}

	ret := &dp[turn][here]
	if *ret != 0 {
		return *ret
	}
	*ret = INF // 최솟값 비교를 위해 INF로 초기화

	for next := 1; next <= N; next++ {
		// here에서 next로 가는 경로가 존재하는 경우
		if path[here][next] != 0 {
			*ret = min(*ret, rec(turn+1, next)+path[here][next]) // 최솟값 갱신
		}
	}

	return *ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
