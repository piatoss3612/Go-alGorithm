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

	R1, C1 int
	R2, C2 int

	move = [8][3][2]int{
		{{0, -1}, {-1, -1}, {-1, -1}},
		{{0, -1}, {+1, -1}, {+1, -1}},

		{{-1, 0}, {-1, -1}, {-1, -1}},
		{{-1, 0}, {-1, +1}, {-1, +1}},

		{{0, +1}, {-1, +1}, {-1, +1}},
		{{0, +1}, {+1, +1}, {+1, +1}},

		{{+1, 0}, {+1, -1}, {+1, -1}},
		{{+1, 0}, {+1, +1}, {+1, +1}},
	}
	visited [10][9]bool
)

// 난이도: Gold 5
// 메모리: 940KB
// 시간: 4ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Input()
	Solve()
}

func Input() {
	R1, C1 = scanInt(), scanInt()
	R2, C2 = scanInt(), scanInt()
}

type Move struct {
	r, c int
	cnt  int
}

func Solve() {
	q := []Move{}
	q = append(q, Move{r: R1, c: C1, cnt: 0})
	visited[R1][C1] = true

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		r, c, cnt := front.r, front.c, front.cnt

		if r == R2 && c == C2 {
			fmt.Fprintln(writer, cnt)
			return
		}

		for i := 0; i < 8; i++ {
			movable := true
			nr, nc := r, c
			for j := 0; j < 3; j++ {
				nr, nc = nr+move[i][j][0], nc+move[i][j][1]
				if !valid(nr, nc) || (j != 2 && nr == R2 && nc == C2) {
					movable = false
					break
				}
			}

			if movable && !visited[nr][nc] {
				visited[nr][nc] = true
				q = append(q, Move{r: nr, c: nc, cnt: cnt + 1})
			}
		}
	}
	fmt.Fprintln(writer, -1)
}

func valid(r, c int) bool {
	return r >= 0 && r < 10 && c >= 0 && c < 9
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
