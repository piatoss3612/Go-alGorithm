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
	a, b    int
	N, M    int
	conn    [][]int
	visited []bool
)

// 14496번: 그대, 그머가 되어
// https://www.acmicpc.net/problem/14496
// 난이도: 실버 2
// 메모리: 1472 KB
// 시간: 8 ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	a, b = scanInt(), scanInt()
	N, M = scanInt(), scanInt()

	conn = make([][]int, N+1)
	visited = make([]bool, N+1)

	for i := 0; i < M; i++ {
		u, v := scanInt(), scanInt()
		conn[u] = append(conn[u], v)
		conn[v] = append(conn[v], u)
	}
}

func Solve() {
	fmt.Fprintln(writer, bfs())
}

func bfs() int {
	q := [][2]int{{a, 0}}
	visited[a] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur[0] == b {
			return cur[1]
		}

		for _, next := range conn[cur[0]] {
			if !visited[next] {
				visited[next] = true
				q = append(q, [2]int{next, cur[1] + 1})
			}
		}
	}

	return -1
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
