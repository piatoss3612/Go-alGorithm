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
	N, M    int
	K       int
	edges   [][]int
	ruins   []int
	ruined  []bool
)

// 18231번: 파괴된 도시
// hhttps://www.acmicpc.net/problem/18231
// 난이도: 골드 5
// 메모리: 5608 KB
// 시간: 40 ms
// 분류: 그래프 이론, 그래프 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	edges = make([][]int, N+1)
	for i := 0; i < M; i++ {
		u, v := scanInt(), scanInt()
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}

	K = scanInt()
	ruins = make([]int, K)
	ruined = make([]bool, N+1)
	for i := 0; i < K; i++ {
		x := scanInt()
		ruins[i] = x
		ruined[x] = true
	}
}

func Solve() {
	bombs := make([]int, 0)
	visited := make([]bool, N+1)

	for _, ruin := range ruins {
		isBomb := true

		for _, neighbor := range edges[ruin] {
			if !ruined[neighbor] {
				isBomb = false
				break
			}
		}

		if isBomb {
			visited[ruin] = true
			for _, neighbor := range edges[ruin] {
				visited[neighbor] = true
			}
			bombs = append(bombs, ruin)
		}
	}

	if len(bombs) == 0 {
		fmt.Fprintln(writer, -1)
		return
	}

	for i := 1; i <= N; i++ {
		if ruined[i] && !visited[i] {
			fmt.Fprintln(writer, -1)
			return
		}
	}

	sort.Ints(bombs)

	fmt.Fprintln(writer, len(bombs))
	for _, bomb := range bombs {
		fmt.Fprint(writer, bomb, " ")
	}
	fmt.Fprintln(writer)
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
