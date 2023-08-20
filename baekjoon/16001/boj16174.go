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
	board   [65][65]int
	visited [65][65]bool
	dy      = [2]int{0, 1}
	dx      = [2]int{1, 0}
)

// 난이도: Silver 1
// 메모리: 948KB
// 시간: 4ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색, 깊이 우선 탐색
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
			board[i][j] = scanInt()
		}
	}
}

func Solve() {
	dfs(1, 1)

	if visited[N][N] {
		fmt.Fprintln(writer, "HaruHaru")
	} else {
		fmt.Fprintln(writer, "Hing")
	}
}

func dfs(y, x int) {
	visited[y][x] = true

	for i := 0; i < 2; i++ {
		ny := y + dy[i]*board[y][x]
		nx := x + dx[i]*board[y][x]

		if ny < 1 || ny > N || nx < 1 || nx > N {
			continue
		}

		if visited[ny][nx] {
			continue
		}

		dfs(ny, nx)
	}
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
