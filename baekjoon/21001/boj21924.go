package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner                = bufio.NewScanner(os.Stdin)
	writer                 = bufio.NewWriter(os.Stdout)
	N, M                   int
	parent                 []int
	roads                  []*Road // 도로 정보
	totalCost, optimalCost int     // totalCost - optimalCost = 모든 건물이 도로를 통해 연결되도록 최소한으로 연결된 도로들의 비용의 합
	edgeCount              int     // 연결된 도로의 개수: 최소 스패닝 트리를 만족하려면 N-1개의 도로가 필요
)

type Road struct {
	a, b int
	cost int
}

// 난이도: Gold 4
// 메모리: 27656KB
// 시간: 412ms
// 분류: 최소 스패닝 트리
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	parent = make([]int, N+1)
	roads = make([]*Road, M)
	for i := 1; i <= N; i++ {
		parent[i] = i
	}

	for i := 0; i < M; i++ {
		roads[i] = &Road{scanInt(), scanInt(), scanInt()}
		totalCost += roads[i].cost
	}

	// 비용이 가장 적게 드는 도로를 기준으로 오름차순 정렬
	sort.Slice(roads, func(i, j int) bool {
		return roads[i].cost < roads[j].cost
	})
}

func Solve() {
	for len(roads) > 0 {
		road := roads[0]
		roads = roads[1:]
		union(road)
	}

	// 도로의 개수가 N-1보다 작은 경우는 모든 건물이 도로를 통해 연결되지 않았음을 의미
	if edgeCount < N-1 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, totalCost-optimalCost)
	}
}

func find(x int) int {
	if x == parent[x] {
		return x
	}

	parent[x] = find(parent[x])
	return parent[x]
}

func union(r *Road) {
	a, b := find(r.a), find(r.b)
	if a != b {
		parent[b] = a
		optimalCost += r.cost
		edgeCount++
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
