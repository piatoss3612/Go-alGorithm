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
	N       int
	color   []int
	tree    [][]int
	visited []bool
)

// 24230번: 트리 색칠하기
// https://www.acmicpc.net/problem/24230
// 난이도: 골드 5
// 메모리: 43740 KB
// 시간: 168 ms
// 분류: 그래프 이론, 그래프 탐색, 트리, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	color = make([]int, N+1)
	for i := 1; i <= N; i++ {
		color[i] = scanInt()
	}
	tree = make([][]int, N+1)
	for i := 1; i < N; i++ {
		u, v := scanInt(), scanInt()
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
	}

	visited = make([]bool, N+1)
}

func Solve() {
	fmt.Fprintln(writer, dfs(1, 0))
}

func dfs(v, c int) (ans int) {
	visited[v] = true
	if color[v] != c {
		ans += 1
	}

	for _, u := range tree[v] {
		if !visited[u] {
			ans += dfs(u, color[v])
		}
	}

	return
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
