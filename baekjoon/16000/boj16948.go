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

	N              int
	r1, c1, r2, c2 int
	board          [200][200]int
	visited        [200][200]bool

	dr = [6]int{-2, -2, 0, 0, 2, 2}
	dc = [6]int{-1, 1, -2, 2, -1, 1}
)

// 난이도: Silver 1
// 메모리: 1740KB
// 시간: 4ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	r1, c1, r2, c2 = scanInt(), scanInt(), scanInt(), scanInt()
}

func Solve() {
	q := [][2]int{{r1, c1}}
	visited[r1][c1] = true

	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]

		if r == r2 && c == c2 {
			fmt.Fprintln(writer, board[r][c])
			return
		}

		for i := 0; i < 6; i++ {
			nr, nc := r+dr[i], c+dc[i]
			if nr < 0 || nr >= N || nc < 0 || nc >= N {
				continue
			}
			if visited[nr][nc] {
				continue
			}

			visited[nr][nc] = true
			board[nr][nc] = board[r][c] + 1
			q = append(q, [2]int{nr, nc})
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
