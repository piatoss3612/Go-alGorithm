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

	w, h    int
	chart   [51][51]int
	visited [51][51]bool

	dy = [8]int{0, 0, 1, -1, 1, 1, -1, -1}
	dx = [8]int{1, -1, 0, 0, 1, -1, 1, -1}
)

// 난이도: Silver 2
// 메모리: 1148KB
// 시간: 8ms
// 분류: 그래프 이론, 그래프 탐색, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	for {
		if !Setup() {
			break
		}
	}
}

func Setup() bool {
	w, h = scanInt(), scanInt()
	if w == 0 && h == 0 {
		return false
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			chart[i][j] = scanInt()
		}
	}

	visited = [51][51]bool{}
	Solve()

	return true
}

func Solve() {
	var ans int
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if chart[i][j] == 1 && !visited[i][j] {
				ans++
				dfs(i, j)
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func dfs(y, x int) {
	visited[y][x] = true

	for i := 0; i < 8; i++ {
		ny, nx := y+dy[i], x+dx[i]
		if inRange(ny, nx) && chart[ny][nx] == 1 && !visited[ny][nx] {
			dfs(ny, nx)
		}
	}
}

func inRange(y, x int) bool {
	return 1 <= y && y <= h && 1 <= x && x <= w
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
