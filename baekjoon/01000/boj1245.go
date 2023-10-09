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

	N, M    int
	farm    [101][71]int
	visited [101][71]bool

	dy = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dx = []int{-1, 0, 1, -1, 1, -1, 0, 1}
)

// 난이도: Gold 5
// 메모리: 1008KB
// 시간: 4ms
// 분류: 그래프 탐색, 깊이 우선 탐색
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
			farm[i][j] = scanInt()
		}
	}
}

func Solve() {
	cnt := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if !visited[i][j] && farm[i][j] != 0 {
				if DFS(i, j) {
					cnt++
				}
			}
		}
	}

	fmt.Fprintln(writer, cnt)
}

func DFS(y, x int) bool {
	visited[y][x] = true
	flag := true

	for i := 0; i < 8; i++ {
		ny, nx := y+dy[i], x+dx[i]
		if valid(ny, nx) {
			// 산봉우리가 될 수 없는 경우
			if farm[y][x] < farm[ny][nx] {
				flag = false
			}

			// 방문하지 않은 같은 높이의 봉우리가 있는 경우
			if !visited[ny][nx] && farm[ny][nx] == farm[y][x] {
				if !DFS(ny, nx) {
					flag = false
				}
			}
		}
	}

	return flag
}

func valid(y, x int) bool {
	return y >= 1 && y <= N && x >= 1 && x <= M
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
