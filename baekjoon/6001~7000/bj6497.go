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
	edges   []Edge
	parent  []int
)

type Edge struct {
	x, y, z int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for {
		m, n := scanInt(), scanInt()

		if m == 0 && n == 0 {
			return
		}

		testCase(m, n)
	}
}

func testCase(m, n int) {
	edges = make([]Edge, n)
	parent = make([]int, m)
	sum := 0

	for i := 0; i < n; i++ {
		edges[i] = Edge{scanInt(), scanInt(), scanInt()}
		sum += edges[i].z
		if i < m {
			parent[i] = i
		}
	}

	// 비용을 오름차순으로 정렬하여 비용이 적은 간선부터 탐색
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].z < edges[j].z
	})

	// 최소 스패닝 트리의 간선의 수는 정점의 수 - 1
	totalEdges := 0

	for _, edge := range edges {
		if totalEdges == m-1 {
			break
		}

		if !sameParent(edge.x, edge.y) {
			union(edge.x, edge.y)
			sum -= edge.z
			totalEdges += 1
		}
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

func sameParent(x, y int) bool {
	x = find(x)
	y = find(y)
	if x == y {
		return true
	}
	return false
}

func union(x, y int) {
	x = find(x)
	y = find(y)
	if x != y {
		parent[y] = x
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
