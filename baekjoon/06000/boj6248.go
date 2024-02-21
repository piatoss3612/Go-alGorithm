package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M, X int
	edges   [][]Edge
)

type Edge struct {
	To, Cost int
}

type PQ []*Edge

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x interface{}) {
	item := x.(*Edge)
	*pq = append(*pq, item)
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

const INF = 987654321

// 6248번: Bronze Cow Party
// hhttps://www.acmicpc.net/problem/6248
// 난이도: 골드 5
// 메모리: 5304 KB
// 시간: 20 ms
// 분류: 데이크스트라, 최단 경로
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, X = scanInt(), scanInt(), scanInt()
	edges = make([][]Edge, N+1)
	for i := 0; i < M; i++ {
		u, v, w := scanInt(), scanInt(), scanInt()
		edges[u] = append(edges[u], Edge{v, w})
		edges[v] = append(edges[v], Edge{u, w})
	}
}

func Solve() {
	fmt.Fprintln(writer, dijkstra()*2)
}

func dijkstra() int {
	dist := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = INF
	}
	dist[X] = 0

	pq := new(PQ)
	heap.Init(pq)
	heap.Push(pq, &Edge{X, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Edge)
		if dist[cur.To] < cur.Cost {
			continue
		}

		for _, next := range edges[cur.To] {
			cost := cur.Cost + next.Cost
			if cost < dist[next.To] {
				dist[next.To] = cost
				heap.Push(pq, &Edge{next.To, cost})
			}
		}
	}

	maxDist := 0
	for i := 1; i <= N; i++ {
		if i == X {
			continue
		}
		maxDist = max(maxDist, dist[i])
	}

	return maxDist
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
