package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	INF     = 987654321
	edges   [][]edge
	n, m, t int
	s, g, h int
)

type edge struct {
	next, cost int
}

// 메모리: 7016KB
// 시간: 68ms
// s에서 x로 가는 최단 경로에 간선 g-h가 포함되어 있는지 확인하는 다익스트라 문제
// 최단 경로가 g-h를 포함하는 것과 포함하지 않는 것까지 여러 개일 수 있으므로 역추적은 사용할 수 없다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T := scanInt() // 테스트 케이스

	for i := 1; i <= T; i++ {
		n, m, t = scanInt(), scanInt(), scanInt()
		edges = make([][]edge, n+1)

		s, g, h = scanInt(), scanInt(), scanInt()

		// 양방향 그래프임에 주의
		var a, b, d int
		for i := 1; i <= m; i++ {
			a, b, d = scanInt(), scanInt(), scanInt()
			edges[a] = append(edges[a], edge{b, d})
			edges[b] = append(edges[b], edge{a, d})
		}

		sDist := dijkstra(s) // s에서 시작하는 최단 거리
		gDist := dijkstra(g) // g에서 시작하는 최단 거리
		hDist := dijkstra(h) // h에서 시작하는 최단 거리

		var x, temp1, temp2 int
		var ans []int
		for i := 1; i <= t; i++ {
			x = scanInt()
			// temp1: s에서 g로 가는 최단 거리 + g에서 h로 가는 최단 거리 + h에서 x로 가는 최단 거리
			// 즉, s에서 x로 가는 최단 경로가 g에서 h로 가는 경로를 포함하고 있는지 확인
			temp1 = sDist[g] + gDist[h] + hDist[x]
			// temp1: s에서 h로 가는 최단 거리 + h에서 g로 가는 최단 거리 + g에서 x로 가는 최단 거리
			// 즉, s에서 x로 가는 최단 경로가 h에서 g로 가는 경로를 포함하고 있는지 확인
			temp2 = sDist[h] + hDist[g] + gDist[x]

			// temp1 또는 temp2가 s에서 x로 가는 최단 거리의 비용과 같은 경우
			if temp1 == sDist[x] || temp2 == sDist[x] {
				ans = append(ans, x)
			}
		}

		sort.Ints(ans)

		for _, v := range ans {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
	}
}

type PQ []edge

func (q PQ) Len() int { return len(q) }
func (q PQ) Less(i, j int) bool {
	return q[i].cost > q[j].cost
}
func (q PQ) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *PQ) Push(x interface{}) {
	*q = append(*q, x.(edge))
}

func (q *PQ) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func dijkstra(v int) []int {
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = INF
	}
	dist[v] = 0

	pq := &PQ{{v, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		top := heap.Pop(pq).(edge)
		next := top.next
		cost := -top.cost

		if dist[next] < cost {
			continue
		}

		for _, eg := range edges[next] {
			if dist[eg.next] > cost+eg.cost {
				dist[eg.next] = cost + eg.cost
				heap.Push(pq, edge{eg.next, -dist[eg.next]})
			}
		}
	}
	return dist
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
