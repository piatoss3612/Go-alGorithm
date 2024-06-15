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
	board   [][]byte
	sy, sx  int
	hit     [][]int
	dy      = []int{0, 0, 1, -1}
	dx      = []int{1, -1, 0, 0}
	q       = [][2]int{}
)

// 31863번: 내진 설계
// hhttps://www.acmicpc.net/problem/31863
// 난이도: 구현, 그래프 이론, 그래프 탐색, 시뮬레이션
// 메모리: 21588 KB
// 시간: 108 ms
// 분류:
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	board = make([][]byte, N)
	hit = make([][]int, N)
	for i := 0; i < N; i++ {
		b := scanBytes()
		board[i] = make([]byte, M)
		hit[i] = make([]int, M)
		for j := 0; j < M; j++ {
			board[i][j] = b[j]
			if board[i][j] == '@' {
				sy, sx = i, j
			}
		}
	}
}

func Solve() {
	fallen := 0

	// @ 진원지 상하좌우 2칸씩
	for i := 0; i < 4; i++ {
		ny, nx := sy+dy[i], sx+dx[i]

		if hitBoard(ny, nx) {
			ny, nx = ny+dy[i], nx+dx[i]

			hitBoard(ny, nx)
		}
	}

	// 나머지 상하좌우 1칸씩
	for len(q) > 0 {
		y, x := q[0][0], q[0][1]
		q = q[1:]
		fallen++

		for i := 0; i < 4; i++ {
			ny, nx := y+dy[i], x+dx[i]

			hitBoard(ny, nx)
		}
	}

	alive := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if board[i][j] == '*' || board[i][j] == '#' {
				alive++
			}
		}
	}

	fmt.Fprintln(writer, fallen, alive)
}

func hitBoard(y, x int) bool {
	if !inRange(y, x) {
		return false
	}

	if board[y][x] == '|' {
		return false
	}

	hit[y][x]++

	if board[y][x] == '*' {
		board[y][x] = '.'
		q = append(q, [2]int{y, x})
	}

	if board[y][x] == '#' && hit[y][x] == 2 {
		board[y][x] = '.'
		q = append(q, [2]int{y, x})
	}

	return true
}

func inRange(y, x int) bool {
	return 0 <= y && y < N && 0 <= x && x < M
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
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
