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
	road    [][]int
	dist    []int
	visited []bool
)

// 난이도: Silver 1
// 메모리: 6472KB
// 시간: 32ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	road = make([][]int, N+1)
	dist = make([]int, N+1)
	visited = make([]bool, N+1)

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		road[a] = append(road[a], b)
		road[b] = append(road[b], a)
	}
}

func Solve() {
	q := [][2]int{{1, 0}}
	visited[1] = true

	maxDist := 0

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		p := front[0]
		d := front[1]
		maxDist = max(maxDist, d)

		for _, next := range road[p] {
			if !visited[next] {
				visited[next] = true
				dist[next] = d + 1
				q = append(q, [2]int{next, d + 1})
			}
		}
	}

	number := N
	cnt := 0
	for i := 1; i <= N; i++ {
		if dist[i] == maxDist {
			cnt++
			number = min(number, i)
		}
	}

	fmt.Fprintln(writer, number, maxDist, cnt)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
