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
	N, M, K int
	board   [101][101]int
	visited [101][101]bool
	count   int
	sum     int
	ans     = -987654321

	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
)

// 18290번: NM과 K (1)
// https://www.acmicpc.net/problem/18290
// 난이도: 실버 1
// 메모리: 880 KB
// 시간: 1520 ms
// 분류: 브루트포스 알고리즘, 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			board[i][j] = scanInt()
		}
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			visit(i, j)
		}
	}

	fmt.Fprintln(writer, ans)
}

func visit(x, y int) {
	count += 1
	visited[x][y] = true
	sum += board[x][y]

	defer func() {
		count -= 1
		visited[x][y] = false
		sum -= board[x][y]
	}()

	if count == K {
		if sum > ans {
			ans = sum
		}
		return
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if !visited[i][j] && visitable(i, j) {
				visit(i, j)
			}
		}
	}
}

func visitable(x, y int) bool {
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if inRange(nx, ny) {
			if visited[nx][ny] {
				return false
			}
		}
	}

	return true
}

func inRange(x, y int) bool {
	return 1 <= x && x <= N && 1 <= y && y <= M
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
