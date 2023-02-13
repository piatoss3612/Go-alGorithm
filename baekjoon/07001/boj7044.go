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
	routes  []Route
)

type Route struct {
	A, B int
	Cost int
}

// 난이도: Gold 4
// 메모리: 1632KB
// 시간: 16ms
// 분류: 최소 스패닝 트리

// (i) the total cost of these connections is as large as possible,
// (ii) all the barns are connected together (so that it is possible to reach any barn from any other barn via a path of installed connections)
// (iii) so that there are no cycles among the connections (which Farmer John would easily be able to detect)
// Conditions (ii) and (iii) ensure that the final set of connections will look like a "tree".
// -> 비용이 최대로 들어가는 최소 스패닝 트리
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

	routes = make([]Route, M)
	for i := 0; i < M; i++ {
		routes[i] = Route{scanInt(), scanInt(), scanInt()}
	}

	// 비용이 가장 큰 연결 정보를 기준으로 내림차순 정렬
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Cost > routes[j].Cost
	})
}

func Solve() {
	edges := 0
	totalCost := 0

	for len(routes) > 0 {
		r := routes[0]
		routes = routes[1:]

		pa, pb := Find(r.A), Find(r.B)
		// A의 부모 요소와 B의 부모 요소가 연결되어 있지 않은 경우
		if pa != pb {
			parent[pb] = pa     // 연결
			totalCost += r.Cost // 비용 추가
			edges++             // 간성의 개수 증가
		}

		// 간선의 개수가 N-1개인 경우 = 최소 스패닝 트리를 구성한 경우
		if edges == N-1 {
			break
		}
	}

	if edges != N-1 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, totalCost)
	}
}

func Find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = Find(parent[x])
	return parent[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
