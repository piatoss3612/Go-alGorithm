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
	dy      = []int{-2, -2, -1, -1, +1, +1, +2, +2}
	dx      = []int{-1, +1, -2, +2, -2, +2, -1, +1}
	T, L    int
	visited [300][300]bool
)

// 난이도: Sliver 1
// 메모리: 7312KB
// 시간: 84ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		TestCase()
	}
}

type Move struct {
	y, x int
	turn int
}

func TestCase() {
	L = scanInt()
	sy, sx := scanInt(), scanInt()
	ty, tx := scanInt(), scanInt()

	visited = [300][300]bool{}

	q := []Move{}
	q = append(q, Move{sy, sx, 0})
	visited[sy][sx] = true

	for len(q) > 0 {
		next := q[0]
		q = q[1:]

		if next.y == ty && next.x == tx {
			fmt.Fprintln(writer, next.turn)
			return
		}

		for i := 0; i < 8; i++ {
			ny, nx := next.y+dy[i], next.x+dx[i]
			if ny >= 0 && ny < L && nx >= 0 && nx < L {
				if !visited[ny][nx] {
					visited[ny][nx] = true
					q = append(q, Move{ny, nx, next.turn + 1})
				}
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
