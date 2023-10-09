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
	dy      = []int{-1, +0, +0, +1}
	dx      = []int{+0, -1, +1, +0}
	N, M, K int
	paths   [][]int
	visited [][]bool
)

// 메모리: 1520KB
// 시간: 8ms
// 분류: 그래프 탐색, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()

	paths = make([][]int, N+1)
	visited = make([][]bool, N+1)

	for i := 1; i <= N; i++ {
		paths[i] = make([]int, M+1)
		visited[i] = make([]bool, M+1)
	}

	K = scanInt()
	for i := 1; i <= K; i++ {
		paths[scanInt()][scanInt()] = 1
	}
}

func Solve() {
	ans := 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if !visited[i][j] {
				if paths[i][j] == 1 {
					ans = max(ans, DFS(i, j))
				} else {
					visited[i][j] = true
				}
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func DFS(y, x int) int {
	visited[y][x] = true
	cnt := 1

	for i := 0; i < 4; i++ {
		ny, nx := y+dy[i], x+dx[i]
		if valid(ny, nx) {
			if !visited[ny][nx] && paths[ny][nx] == 1 {
				cnt += DFS(ny, nx)
			}
		}
	}
	return cnt
}

func valid(y, x int) bool {
	if y >= 1 && y <= N && x >= 1 && x <= M {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
