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
	room    [51][51]int
	dx      = [4]int{0, 0, 1, -1}
	dy      = [4]int{1, -1, 0, 0}
)

// 23352번: 방탈출
// hhttps://www.acmicpc.net/problem/23352
// 난이도: 골드 5
// 메모리: 35952 KB
// 시간: 292 ms
// 분류: 너비 우선 탐색, 브루트포스 알고리즘
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
			room[i][j] = scanInt()
		}
	}
}

func Solve() {
	maxMove := 0
	maxSum := 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if room[i][j] != 0 {
				move, sum := BFS(i, j)
				if move > maxMove {
					maxMove = move
					maxSum = sum
				} else if move == maxMove {
					maxSum = max(maxSum, sum)
				}
			}
		}
	}

	fmt.Fprintln(writer, maxSum)
}

type Move struct {
	x, y int
	cnt  int
}

func BFS(sx, sy int) (maxMove int, maxSum int) {
	visited := [51][51]bool{}
	q := make([]Move, 0)
	q = append(q, Move{sx, sy, 0})
	visited[sx][sy] = true

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		x, y, cnt := front.x, front.y, front.cnt

		if cnt > maxMove {
			maxMove = cnt
			maxSum = room[sx][sy] + room[x][y]
		} else if cnt == maxMove {
			maxSum = max(maxSum, room[sx][sy]+room[x][y])
		}

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if inRange(nx, ny) && !visited[nx][ny] && room[nx][ny] != 0 {
				visited[nx][ny] = true
				q = append(q, Move{nx, ny, cnt + 1})
			}
		}
	}

	return
}

func inRange(x, y int) bool {
	return 1 <= x && x <= N && 1 <= y && y <= M
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
