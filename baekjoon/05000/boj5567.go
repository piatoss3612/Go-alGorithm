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

	N, M       int
	areFriends [501][501]bool
	visited    [501]bool
)

// 난이도: Silver 2
// 메모리: 1284KB
// 시간: 12ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		// a, b가 친구라면 b, a도 친구
		areFriends[a][b] = true
		areFriends[b][a] = true
	}
}

func Solve() {
	fmt.Fprintln(writer, bfs())
}

func bfs() int {
	q := [][2]int{{1, 0}}
	visited[1] = true

	cnt := 0

	for len(q) > 0 {
		here, depth := q[0][0], q[0][1]
		q = q[1:]

		// 친구의 친구까지만 탐색
		if depth == 2 {
			continue
		}

		for there := 1; there <= N; there++ {
			if !visited[there] && areFriends[here][there] {
				q = append(q, [2]int{there, depth + 1})
				visited[there] = true
				cnt++
			}
		}
	}

	return cnt
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
