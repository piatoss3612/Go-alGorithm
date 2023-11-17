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
	screen  [1001][1001]int
	T       int
	visited [1001][1001]bool
)

// 21938번: 영상처리
// https://www.acmicpc.net/problem/21938
// 난이도: 실버 2
// 메모리: 15132 KB
// 시간: 164 ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색, 깊이 우선 탐색
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
			r, g, b := scanInt(), scanInt(), scanInt()
			avg := (r + g + b) / 3
			screen[i][j] = avg
		}
	}
	T = scanInt()
}

func Solve() {
	ans := 0
	q := [][2]int{}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if screen[i][j] >= T && !visited[i][j] {
				q = append(q, [2]int{i, j})
				visited[i][j] = true
				ans++

				for len(q) > 0 {
					x, y := q[0][0], q[0][1]
					q = q[1:]

					for _, d := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
						nx, ny := x+d[0], y+d[1]
						if nx < 1 || ny < 1 || nx > N || ny > M {
							continue
						}
						if screen[nx][ny] >= T && !visited[nx][ny] {
							q = append(q, [2]int{nx, ny})
							visited[nx][ny] = true
						}
					}
				}
			}
		}
	}

	fmt.Fprintln(writer, ans)
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
