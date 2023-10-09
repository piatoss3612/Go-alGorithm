package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	M, N, K int
	visited [100][100]int
)

// 난이도: Silver 1
// 메모리: 7348KB
// 시간: 12ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	M, N, K = scanInt(), scanInt(), scanInt()
	for i := 0; i < K; i++ {
		x1, y1, x2, y2 := scanInt(), scanInt(), scanInt(), scanInt()
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				visited[y][x] = 1
			}
		}
	}
}

func Solve() {
	ans := []int{}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if visited[i][j] == 0 {
				ans = append(ans, dfs(i, j))
			}
		}
	}

	sort.Ints(ans)

	fmt.Fprintf(writer, "%d\n", len(ans))

	for _, v := range ans {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func dfs(y, x int) int {
	visited[y][x] = 1

	cnt := 1

	for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		ny, nx := y+d[0], x+d[1]
		if inRange(ny, nx) && visited[ny][nx] == 0 {
			cnt += dfs(ny, nx)
		}
	}

	return cnt
}

func inRange(y, x int) bool {
	return 0 <= y && y < M && 0 <= x && x < N
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
