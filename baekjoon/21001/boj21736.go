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
	campus  [][]byte
	visited [][]bool
	dy      = [4]int{0, 0, 1, -1}
	dx      = [4]int{1, -1, 0, 0}
)

// 난이도: Silver 2
// 메모리: 10996KB
// 시간: 44ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	campus = make([][]byte, N)
	visited = make([][]bool, N)
	for i := 0; i < N; i++ {
		campus[i] = []byte(scanString())
		visited[i] = make([]bool, M)
	}
}

func Solve() {
	var y, x int
Loop:
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if campus[i][j] == 'I' {
				y, x = i, j
				break Loop
			}
		}
	}

	q := [][2]int{{y, x}}
	visited[y][x] = true
	cnt := 0

	for len(q) > 0 {
		y, x = q[0][0], q[0][1]
		q = q[1:]

		if campus[y][x] == 'P' {
			cnt++
		}

		for i := 0; i < 4; i++ {
			ny, nx := y+dy[i], x+dx[i]
			if valid(ny, nx) && !visited[ny][nx] && campus[ny][nx] != 'X' {
				visited[ny][nx] = true
				q = append(q, [2]int{ny, nx})
			}
		}
	}

	if cnt == 0 {
		fmt.Fprintln(writer, "TT")
	} else {
		fmt.Fprintln(writer, cnt)
	}
}

func valid(y, x int) bool {
	return 0 <= y && y < N && 0 <= x && x < M
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
