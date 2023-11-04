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
	u, v, t int
}

// 27945번: 슬슬 가지를 먹지 않으면 죽는다
// https://www.acmicpc.net/problem/27945
// 난이도: 골드 3
// 메모리: 13916 KB
// 시간: 288 ms
// 분류: 자료 구조, 그래프 이론, 그리디 알고리즘, 분리 집합, 최소 스패닝 트리
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
		edges[i] = Edge{scanInt(), scanInt(), scanInt()}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].t < edges[j].t
	})
}

func Solve() {
	cnt := 0

	for len(edges) > 0 && cnt < N-1 {
		edge := edges[0]
		edges = edges[1:]

		pu, pv := find(edge.u), find(edge.v)
		if pu != pv {
			if edge.t != cnt+1 {
				break
			}
			parent[pu] = pv
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt+1)
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
