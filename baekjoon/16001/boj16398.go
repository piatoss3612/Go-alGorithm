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
	parent  []int
)

type edge struct {
	a, b int
	cost int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	edges := []edge{}
	parent = make([]int, n+1)

	// 그래프가 i = j인 대각선을 기준으로 대칭 형태이므로 모든 입력값을 저장할 필요가 없다
	var cost int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			cost = scanInt()
			if j > i {
				edges = append(edges, edge{i, j, cost})
			}
		}
		parent[i] = i
	}

	// 비용을 기준으로 오름차순 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	totalEdges := 0
	minCost := 0

	// 크루스칼 알고리즘
	for _, edge := range edges {
		if totalEdges == n-1 {
			break
		}

		if !sameParent(edge.a, edge.b) {
			minCost += edge.cost
			totalEdges += 1
			union(edge.a, edge.b)
		}
	}

	fmt.Fprintln(writer, minCost)
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	x = find(x)
	y = find(y)
	if x != y {
		parent[y] = x
	}
}

func sameParent(x, y int) bool {
	x = find(x)
	y = find(y)
	if x == y {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
