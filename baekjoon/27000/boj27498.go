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
	parent  []int
	edges   []Edge
)

type Edge struct {
	a, b, c, d int
}

// 27498번: 연애 혁명
// https://www.acmicpc.net/problem/27498
// 난이도: 골드 3
// 메모리: 4188 KB
// 시간: 60 ms
// 분류: 그래프 이론, 최소 스패닝 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	parent = make([]int, N+1)
	for i := 1; i <= N; i++ {
		parent[i] = i
	}

	edges = make([]Edge, M)
	for i := 0; i < M; i++ {
		edges[i] = Edge{scanInt(), scanInt(), scanInt(), scanInt()}
		if edges[i].d == 1 {
			edges[i].c += 10000
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].c > edges[j].c
	})
}

func Solve() {
	cnt := 0
	sum := 0

	for len(edges) > 0 && cnt < N-1 {
		edge := edges[0]
		edges = edges[1:]

		pa, pb := find(edge.a), find(edge.b)

		if pa != pb {
			parent[pb] = pa
			cnt++
		} else {
			sum += edge.c
		}
	}

	for i := 0; i < len(edges); i++ {
		sum += edges[i].c
	}

	fmt.Fprintln(writer, sum)
}

func find(x int) int {
	if parent[x] == x {
		return x
	}

	parent[x] = find(parent[x])
	return parent[x]
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
