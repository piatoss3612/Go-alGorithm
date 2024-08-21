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
	knight  [2]int
	enemies [][2]int
	board   [501][501]int
	dx      = [8]int{1, 2, 2, 1, -1, -2, -2, -1}
	dy      = [8]int{2, 1, -1, -2, -2, -1, 1, 2}
)

// 18404번: 현명한 나이트
// https://www.acmicpc.net/problem/18404
// 난이도: 실버 1
// 메모리: 13512 KB
// 시간: 20 ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	knight[0], knight[1] = scanInt(), scanInt()
	enemies = make([][2]int, M+1)
	for i := 1; i <= M; i++ {
		x, y := scanInt(), scanInt()
		enemies[i] = [2]int{x, y}
		board[x][y] = 1
	}
}

func Solve() {
	visited := [501][501]int{}
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{knight[0], knight[1]})

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		cnt := visited[x][y]
		queue = queue[1:]

		for i := 0; i < 8; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if !inRange(nx, ny) || visited[nx][ny] != 0 {
				continue
			}

			visited[nx][ny] = cnt + 1
			queue = append(queue, [2]int{nx, ny})
		}
	}

	for i := 1; i <= M; i++ {
		fmt.Fprintf(writer, "%d ", visited[enemies[i][0]][enemies[i][1]])
	}
	fmt.Fprintln(writer)
}

func inRange(x, y int) bool {
	return 1 <= x && x <= N && 1 <= y && y <= N
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
