package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N, M  int
	edges [][]Edge
	dist  [501]int
)

type Edge struct {
	to   int
	cost int
}

const INF = 987654321

// 11657번: 타임머신
// hhttps://www.acmicpc.net/problem/11657
// 난이도: 골드 4
// 메모리: 1268 KB
// 시간: 16 ms
// 분류: 벨만-포드, 최단 경로
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	edges = make([][]Edge, N+1)
	for i := 0; i < M; i++ {
		u, v, c := scanInt(), scanInt(), scanInt()
		edges[u] = append(edges[u], Edge{v, c})
	}
}

func Solve() {
	// 음수 간선이 존재하므로 벨만-포드 알고리즘을 사용한다.

	// dist[i]는 1번 정점에서 i번 정점까지의 최단 거리를 저장한다.
	for i := 2; i <= N; i++ {
		dist[i] = INF
	}

	// dist[1] = 0

	// 음수 사이클이 존재한다면 -1을 출력한다.
	if hasNegativeCycle() {
		fmt.Fprintln(writer, -1)
		return
	}

	// 1번 정점에서 각 정점까지의 최단 거리를 출력한다.
	for i := 2; i <= N; i++ {
		// INF는 1번 정점에서 i번 정점까지의 경로가 없다는 의미이다.
		if dist[i] == INF {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintln(writer, dist[i])
		}
	}
}

func hasNegativeCycle() bool {
	// N-1번 반복한다.
	for i := 1; i <= N-1; i++ {
		// 갱신이 일어나지 않는다면 최단 거리가 갱신되지 않았다는 의미 + 음수 사이클이 존재하지 않는다.
		if !relax() {
			return false
		}
	}

	// N번째에도 갱신이 일어난다면 음수 사이클이 존재한다.
	return relax()
}

func relax() bool {
	updated := false

	// 모든 정점에 대해 갱신을 시도한다.
	for from := 1; from <= N; from++ {
		for _, e := range edges[from] {
			// 1번 정점과 연결되어 있지 않다면 갱신을 시도하지 않는다.
			if dist[from] == INF {
				continue
			}

			to, cost := e.to, e.cost
			// 최단 거리를 갱신한다.
			if dist[from]+cost < dist[to] {
				dist[to] = dist[from] + cost
				updated = true
			}
		}
	}

	return updated
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
