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
	planet  [1000][1000]int
	visited [1000][1000]bool
	dy      = [4]int{0, 0, 1, -1}
	dx      = [4]int{1, -1, 0, 0}
)

// 난이도: Gold 5
// 메모리: 115436KB
// 시간: 324ms
// 분류: 그래프 이론, 그래프 탐색, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			planet[i][j] = scanInt()
		}
	}
}

func Solve() {
	cnt := 0

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if !visited[i][j] && planet[i][j] == 0 {
				dfs(i, j)
				cnt++
			}
		}
	}

	fmt.Fprintln(writer, cnt)
}

func dfs(i, j int) {
	visited[i][j] = true

	for k := 0; k < 4; k++ {
		ny, nx := (i+dy[k]+N)%N, (j+dx[k]+M)%M // 행성은 연결되어 있으므로, 행과 열의 범위를 벗어나면 반대쪽으로 이동할 수 있도록 한다.

		if !visited[ny][nx] && planet[ny][nx] == 0 {
			dfs(ny, nx)
		}
	}
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
