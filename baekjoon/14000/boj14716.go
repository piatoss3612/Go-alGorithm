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
	M, N    int
	dy      = []int{-1, -1, +0, +1, +1, +1, +0, -1}
	dx      = []int{+0, +1, +1, +1, +0, -1, -1, -1}
	banner  [251][251]int
	visited [251][251]bool
)

// 난이도: Silver 1
// 메모리: 8108KB
// 시간: 20ms
// 분류: 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	M, N = scanInt(), scanInt()
	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			banner[i][j] = scanInt()
		}
	}
}

func Solve() {
	ans := 0
	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			if banner[i][j] == 0 {
				visited[i][j] = true
			} else if banner[i][j] == 1 && !visited[i][j] {
				ans++
				DFS(i, j)
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func DFS(y, x int) {
	visited[y][x] = true

	for i := 0; i < 8; i++ {
		ny, nx := y+dy[i], x+dx[i]
		if valid(ny, nx) && banner[ny][nx] == 1 && !visited[ny][nx] {
			DFS(ny, nx)
		}
	}
}

func valid(y, x int) bool {
	if y >= 1 && y <= M && x >= 1 && x <= N {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
