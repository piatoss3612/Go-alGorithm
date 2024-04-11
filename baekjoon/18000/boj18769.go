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
	T       int
	R, C    int
	edges   []Edge
	parent  []int
)

type Edge struct {
	from, to int
	cost     int
}

// 18769번: 그리드 네트워크
// hhttps://www.acmicpc.net/problem/18769
// 난이도: 골드 4
// 메모리: 60680 KB
// 시간: 424 ms
// 분류: 최소 스패닝 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
}

func Setup() {
	T = scanInt()
	for t := 0; t < T; t++ {
		R, C = scanInt(), scanInt()
		edges = make([]Edge, 0, R*(C-1)+(R-1)*C)
		parent = make([]int, R*C+1)

		for i := 1; i <= R*C; i++ {
			parent[i] = i
		}

		for i := 1; i <= R; i++ {
			for j := 1; j <= C-1; j++ {
				edges = append(edges, Edge{boxNum(i, j), boxNum(i, j+1), scanInt()})
			}
		}

		for i := 1; i <= R-1; i++ {
			for j := 1; j <= C; j++ {
				edges = append(edges, Edge{boxNum(i, j), boxNum(i+1, j), scanInt()})
			}
		}

		sort.Slice(edges, func(i, j int) bool {
			return edges[i].cost < edges[j].cost
		})

		Solve()
	}
}

func Solve() {
	cnt := 0
	total := 0

	for _, edge := range edges {
		if isSameSet(edge.from, edge.to) {
			continue
		}

		union(edge.from, edge.to)
		total += edge.cost
		cnt++

		if cnt == R*C-1 {
			break
		}
	}

	fmt.Fprintln(writer, total)
}

func find(x int) int {
	if x == parent[x] {
		return x
	}

	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		parent[y] = x
	}
}

func isSameSet(x, y int) bool {
	return find(x) == find(y)
}

func boxNum(i, j int) int {
	return (i-1)*C + j
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
