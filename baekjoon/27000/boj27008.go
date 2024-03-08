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
	F, P, C, M int
	edges      [][]Edge
)

const INF = 9876543210

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

// 27008번: Checking an Alibi
// hhttps://www.acmicpc.net/problem/27008
// 난이도: 골드 5
// 메모리: 1052 KB
// 시간: 4 ms
// 분류: 데이크스트라
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	F, P, C, M = scanInt(), scanInt(), scanInt(), scanInt()
	edges = make([][]Edge, F+1)
	for i := 1; i <= P; i++ {
		from, to, cost := scanInt(), scanInt(), scanInt()
		edges[from] = append(edges[from], Edge{to, cost})
		edges[to] = append(edges[to], Edge{from, cost})
	}

}

func Solve() {
	dist := dijkstra(1)

	suspects := []int{}

	for i := 1; i <= C; i++ {
		pos := scanInt()
		if dist[pos] <= M {
			suspects = append(suspects, i)
		}
	}

	fmt.Fprintln(writer, len(suspects))
	for _, s := range suspects {
		fmt.Fprintln(writer, s)
	}
}

func dijkstra(start int) (dist []int) {
	dist = make([]int, F+1)
	for i := 1; i <= F; i++ {
		dist[i] = INF
	}
	dist[start] = 0

	pq := new(PQ)
	heap.Init(pq)
	heap.Push(pq, &Edge{start, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Edge)
		if dist[cur.to] < cur.cost {
			continue
		}

		for _, next := range edges[cur.to] {
			if dist[next.to] > cur.cost+next.cost {
				dist[next.to] = cur.cost + next.cost
				heap.Push(pq, &Edge{next.to, dist[next.to]})
			}
		}
	}
	return
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
