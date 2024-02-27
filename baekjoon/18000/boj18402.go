package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, E, T, M int
	edges      [][]Edge
)

const INF = 987654321

type Edge struct {
	to, cost int
}

type PQ []*Edge

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
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

// 18402번: RUN
// hhttps://www.acmicpc.net/problem/18402
// 난이도: 골드 5
// 메모리: 876 KB
// 시간: 8 ms
// 분류: 데이크스트라, 최단 경로
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, E, T, M = scanInt(), scanInt(), scanInt(), scanInt()
	edges = make([][]Edge, N+1)

	for i := 0; i < M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		// 단방향 그래프
		// exit -> 다른 모든 지점으로 가는 최단 경로를 구해야 하므로 역방향으로 저장
		edges[b] = append(edges[b], Edge{to: a, cost: c})
	}
}

func Solve() {
	if E > N {
		fmt.Fprintln(writer, 0)
		return
	}

	dist := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = INF
	}
	dist[E] = 0

	pq := new(PQ)
	heap.Init(pq)
	heap.Push(pq, &Edge{to: E, cost: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Edge)
		if dist[cur.to] < cur.cost {
			continue
		}

		for _, next := range edges[cur.to] {
			if next.cost+cur.cost < dist[next.to] {
				dist[next.to] = next.cost + cur.cost
				heap.Push(pq, &Edge{to: next.to, cost: dist[next.to]})
			}
		}
	}

	cnt := 0

	for i := 1; i <= N; i++ {
		if dist[i] <= T {
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
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
