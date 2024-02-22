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
	visible []int
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

const INF = 987654321321

// 17396번: 백도어
// hhttps://www.acmicpc.net/problem/17396
// 난이도: 골드 5
// 메모리: 31572 KB
// 시간: 160 ms
// 분류: 데이크스트라, 최단 경로
// 주의: INF를 충분히 큰 값으로 설정해야 함
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	visible = make([]int, N)
	edges = make([][]Edge, N)
	for i := 0; i < N; i++ {
		visible[i] = scanInt()
	}
	for i := 0; i < M; i++ {
		a, b, t := scanInt(), scanInt(), scanInt()
		edges[a] = append(edges[a], Edge{b, t})
		edges[b] = append(edges[b], Edge{a, t})
	}
}

func Solve() {
	fmt.Fprintln(writer, dijkstra())
}

func dijkstra() int {
	dist := make([]int, N)
	for i := 1; i < N; i++ {
		dist[i] = INF
	}

	pq := new(PQ)
	heap.Init(pq)
	heap.Push(pq, &Edge{0, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Edge)
		if dist[cur.To] < cur.Cost {
			continue
		}

		for _, next := range edges[cur.To] {
			cost := cur.Cost + next.Cost
			if visible[next.To] == 1 && next.To != N-1 {
				continue
			}

			if cost < dist[next.To] {
				dist[next.To] = cost
				heap.Push(pq, &Edge{next.To, cost})
			}
		}
	}

	if dist[N-1] == INF {
		return -1
	}

	return dist[N-1]
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
