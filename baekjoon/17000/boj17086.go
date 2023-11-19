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
	area    [51][51]int

	dy = [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dx = [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
)

// 17086번: 아기 상어 2
// https://www.acmicpc.net/problem/17086
// 난이도: 실버 2
// 메모리: 15840 KB
// 시간: 168 ms
// 분류: 그래프 이론, 브루트포스 알고리즘, 그래프 탐색, 너비 우선 탐색
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
			area[i][j] = scanInt()
		}
	}
}

func Solve() {
	ans := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if area[i][j] == 0 {
				ans = max(ans, bfs(i, j))
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func bfs(y, x int) int {
	visited := [51][51]bool{}
	q := [][3]int{{y, x, 0}}
	visited[y][x] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for i := 0; i < 8; i++ {
			ny, nx := cur[0]+dy[i], cur[1]+dx[i]
			if inRange(ny, nx) && !visited[ny][nx] {
				if area[ny][nx] == 1 {
					return cur[2] + 1
				}
				visited[ny][nx] = true
				q = append(q, [3]int{ny, nx, cur[2] + 1})
			}
		}
	}

	return 0
}

func inRange(y, x int) bool {
	return 1 <= y && y <= N && 1 <= x && x <= M
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
