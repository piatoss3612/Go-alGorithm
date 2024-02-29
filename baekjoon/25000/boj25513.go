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
	board   [5][5]int
	visited [7][5][5]bool
	r, c    int
	dy      = [4]int{0, 0, 1, -1}
	dx      = [4]int{1, -1, 0, 0}
)

// 25513번: 빠른 오름차순 숫자 탐색
// hhttps://www.acmicpc.net/problem/25513
// 난이도: 골드 5
// 메모리: 900 KB
// 시간: 4 ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			board[i][j] = scanInt()
		}
	}

	r, c = scanInt(), scanInt()
}

type Move struct {
	h, y, x, cnt int
}

func Solve() {
	queue := make([]Move, 0)
	queue = append(queue, Move{0, r, c, 0})
	visited[0][r][c] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.h == 6 {
			fmt.Fprintln(writer, cur.cnt)
			return
		}

		for i := 0; i < 4; i++ {
			ny, nx := cur.y+dy[i], cur.x+dx[i]
			if ny < 0 || ny >= 5 || nx < 0 || nx >= 5 || visited[cur.h][ny][nx] || board[ny][nx] == -1 {
				continue
			}

			visited[cur.h][ny][nx] = true
			if board[ny][nx] == cur.h+1 {
				visited[cur.h+1][ny][nx] = true
				queue = append(queue, Move{cur.h + 1, ny, nx, cur.cnt + 1})
			} else {
				queue = append(queue, Move{cur.h, ny, nx, cur.cnt + 1})
			}
		}
	}

	fmt.Fprintln(writer, -1)
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
