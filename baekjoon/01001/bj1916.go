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
	INF     = 987654321
	edges   [1001][]edge
	dist    [1001]int
	n, m    int
)

type edge struct {
	next, cost int
}

// 메모리: 15880KB
// 시간: 92ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	var a, b, c int
	for i := 1; i <= m; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()
		edges[a] = append(edges[a], edge{b, c})
	}

	s, e := scanInt(), scanInt()

	dijkstra(s)

	fmt.Fprintln(writer, dist[e])
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

func dijkstra(v int) {
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
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
