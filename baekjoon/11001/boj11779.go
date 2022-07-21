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
	nodes   [1001]int // s부터 i번째 노드까지의 최소 비용을 구했을 때 i번 노드로 가기 이전의 노드를 저장
	n, m    int
)

type edge struct {
	next, cost int
}

// 메모리: 15888KB
// 시간: 88ms
// 1916번 문제 + 최소 비용을 구할 수 있는 경로를 역추적하는 문제
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

	var track []int

	// 최소 비용 경로 역추적
	for prev := e; prev != 0; {
		track = append(track, prev)
		prev = nodes[prev]
	}

	fmt.Fprintln(writer, len(track))
	for i := len(track) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%d ", track[i])
	}
	fmt.Fprintln(writer)
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
			// 최소 비용 경로를 찾은 경우
			if dist[eg.next] > cost+eg.cost {
				dist[eg.next] = cost + eg.cost
				heap.Push(pq, edge{eg.next, -dist[eg.next]})
				nodes[eg.next] = next // 최소 비용 경로에서 eg.next의 이전 노드가 next라는 것을 저장
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
