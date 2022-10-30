package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M    int
	locs    [1001]Location // 우주신의 위치
	parent  [1001]int      // 연결 관계의 루트 요소
	edges   []Edge         // 우주신들의 연결 정보
	ans     float64        // 최소 통로 길이
)

type Location struct {
	x, y int
}

type Edge struct {
	a, b     int
	distance float64
}

// 난이도: Gold 3
// 메모리: 60704KB
// 시간: 268ms
// 분류: 최소 스패닝 트리
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		locs[i] = Location{scanInt(), scanInt()}
		parent[i] = i
	}

	for i := 1; i <= M; i++ {
		union(scanInt(), scanInt())
	}
}

func Solve() {
	// 아직 연결되지 않은 우주신들 사이의 연결 정보 전처리
	for i := 1; i <= N-1; i++ {
		for j := i + 1; j <= N; j++ {
			if find(i) != find(j) {
				edges = append(edges, Edge{i, j, distanceAtoB(i, j)})
			}
		}
	}

	// 연결 거리가 가장 짧은 간선을 기준으로 오름차순 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	for len(edges) > 0 {
		e := edges[0]
		edges = edges[1:]

		if union(e.a, e.b) {
			ans += e.distance
		}
	}

	fmt.Fprintf(writer, "%0.2f\n", ans)
}

func find(x int) int {
	if x == parent[x] {
		return x
	}

	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) bool {
	x, y = find(x), find(y)
	if x != y {
		parent[y] = x
		return true
	}
	return false
}

func distanceAtoB(a, b int) float64 {
	return math.Sqrt(math.Pow(float64(locs[a].x-locs[b].x), 2) + math.Pow(float64(locs[a].y-locs[b].y), 2))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
