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
	field   [101][101]int
	visited [101][101]bool

	dy = [4]int{0, 0, 1, -1}
	dx = [4]int{1, -1, 0, 0}
)

// 난이도: Silver 1
// 메모리: 1020KB
// 시간: 8ms
// 분류: 깊이 우선 탐색(DFS)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= M; i++ {
		line := scanBytes()

		for j := 1; j <= N; j++ {
			if line[j-1] == 'W' {
				field[i][j] = 1
			} else {
				field[i][j] = 2
			}
		}
	}
}

func Solve() {
	w, b := 0, 0

	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			if !visited[i][j] {
				temp := dfs(i, j, field[i][j])
				if field[i][j] == 1 {
					w += temp * temp
				} else {
					b += temp * temp
				}
			}
		}
	}

	fmt.Fprintln(writer, w, b)
}

func dfs(y, x, color int) int {
	visited[y][x] = true
	cnt := 1

	for i := 0; i < 4; i++ {
		ny, nx := y+dy[i], x+dx[i]
		if isValid(ny, nx) && !visited[ny][nx] && field[ny][nx] == color {
			cnt += dfs(ny, nx, color)
		}
	}

	return cnt
}

func isValid(y, x int) bool {
	return y >= 1 && y <= M && x >= 1 && x <= N
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
