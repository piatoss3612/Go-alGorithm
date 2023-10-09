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
	N       int
	cows    [][2]int // cows[i]: i번 소의 x, y 좌표
	parent  []int    // parent[i]: i번 소가 속한 집합의 대표
	edges   []Edge   // edges[i]: a번 소와 b번 소 사이의 거리의 제곱 c
)

type Edge struct {
	a, b, c int
}

// 난이도: Gold 3
// 메모리: 13812KB
// 시간: 160ms
// 분류: 최소 스패닝 트리
// 풀이: 모든 소들을 연결하는 최소 스패닝 트리를 구하고, 최소 비용(X: 최소 스패닝 트리를 구성하는 간선 중 가장 긴 간선의 길이의 제곱)을 출력한다.
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	cows = make([][2]int, N+1)
	parent = make([]int, N+1)
	edges = make([]Edge, 0, N*N)

	for i := 1; i <= N; i++ {
		x, y := scanInt(), scanInt()
		cows[i] = [2]int{x, y}
		parent[i] = i
	}

	for i := 1; i <= N-1; i++ {
		for j := i + 1; j <= N; j++ {
			edges = append(edges, Edge{i, j, dist(i, j)}) // 결국은 최소비용 X를 구하는 문제이므로 간선의 길이를 제곱한 값을 저장한다.
		}
	}

	// 간선의 길이를 기준으로 오름차순 정렬한다.
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].c < edges[j].c
	})
}

func Solve() {
	cnt := 0 // 현재까지 연결된 간선의 개수
	amt := 0 // 현재까지 연결된 간선 중 가장 긴 간선의 길이의 제곱

	// 모든 소들을 연결하는 최소 스패닝 트리를 구한다.
	// 연결된 간선의 개수가 N-1개가 되면 종료한다.
	for len(edges) > 0 && cnt < N-1 {
		e := edges[0]
		edges = edges[1:]

		pa, pb := find(e.a), find(e.b)

		// union 연산
		if pa != pb {
			parent[pb] = pa
			cnt += 1
			amt = max(amt, e.c)
		}
	}

	fmt.Fprintln(writer, amt)
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func dist(a, b int) int {
	x, y := cows[a][0]-cows[b][0], cows[a][1]-cows[b][1]
	return x*x + y*y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
