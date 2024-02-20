package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner      = bufio.NewScanner(os.Stdin)
	writer       = bufio.NewWriter(os.Stdout)
	T, C, Ts, Te int
	edges        [][]Edge
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

// 5996번: Heat Wave
// hhttps://www.acmicpc.net/problem/5996
// 난이도: 골드 5
// 메모리: 1564 KB
// 시간: 8 ms
// 분류: 데이크스트라, 최단 경로
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	T, C, Ts, Te = scanInt(), scanInt(), scanInt(), scanInt()
	edges = make([][]Edge, T+1)
	for i := 0; i < C; i++ {
		f, t, c := scanInt(), scanInt(), scanInt()
		// bidirectional
		edges[f] = append(edges[f], Edge{t, c})
		edges[t] = append(edges[t], Edge{f, c})
	}
}

func Solve() {
	fmt.Fprintln(writer, dijkstra())
}

func dijkstra() int {
	dist := make([]int, T+1)
	for i := 1; i <= T; i++ {
		dist[i] = INF
	}
	dist[Ts] = 0

	pq := new(PQ)
	heap.Init(pq)
	heap.Push(pq, &Edge{Ts, 0})

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

	return dist[Te]
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
