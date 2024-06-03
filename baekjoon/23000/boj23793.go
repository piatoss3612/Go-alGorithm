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
	N, M    int
	X, Y, Z int
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

const INF = 9876543210

// 23793번: 두 단계 최단 경로 1
// hhttps://www.acmicpc.net/problem/23793
// 난이도: Gold 4
// 메모리: 20188 KB
// 시간: 216 ms
// 분류: 데이크스트라, 최단 경로
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
		u, v, w := scanInt(), scanInt(), scanInt()
		// uni-directional
		edges[u] = append(edges[u], Edge{v, w})
	}

	X, Y, Z = scanInt(), scanInt(), scanInt()
}

func Solve() {
	xToY := Dijkstra(X, Y, false)
	yToZ := Dijkstra(Y, Z, false)
	xToZ := Dijkstra(X, Z, true)

	var a, b int

	if xToY == INF || yToZ == INF {
		a = -1
	} else {
		a = xToY + yToZ
	}

	if xToZ == INF {
		b = -1
	} else {
		b = xToZ
	}

	fmt.Fprintln(writer, a, b)
}

func Dijkstra(start, end int, exceptY bool) int {
	dist := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = INF
	}

	pq := new(PQ)
	heap.Init(pq)

	dist[start] = 0
	heap.Push(pq, &Edge{start, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Edge)

		if dist[cur.To] < cur.Cost {
			continue
		}

		for _, next := range edges[cur.To] {
			// exceptY가 true이면 Y로 가는 경로는 제외
			if exceptY && (next.To == Y) {
				continue
			}

			nextCost := cur.Cost + next.Cost

			if dist[next.To] > nextCost {
				dist[next.To] = nextCost
				heap.Push(pq, &Edge{next.To, nextCost})
			}
		}
	}

	return dist[end]
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
