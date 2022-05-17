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
	minEdge [1001]int
	n, m    int
)

type edge struct {
	next, cost int
}

// 메모리: 13352KB
// 시간: 60ms
// 연결된 모든 회선이 쌍방향인 컴퓨터들을 최소 개수의 회선으로 복구하는 문제
// 즉, 임의의 노드에서 다른 모든 노드로 가는 최소 비용을 구하는 다익스트라 문제

// 최소 비용이 여러 가지더라도 결과적으로 최소 개수의 회선을 구하면 되므로
// 회선의 연결 정보는 역추적을 사용해 구하면 된다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	var a, b, c int
	for i := 1; i <= m; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()
		edges[a] = append(edges[a], edge{b, c}) // a->b
		edges[b] = append(edges[b], edge{a, c}) // b->a
	}

	// 컴퓨터들이 쌍방향으로 연결되어 있으므로 최소 경로 탐색을 어디에서 시작하든 상관없지만 편의를 위해 1에서 시작
	dijkstra(1)

	type line struct { // 회선 정보를 저장하는 구조체
		s, e int
	}

	check := make(map[line]bool) // 맵을 사용해 회선 정보가 중복되는지 확인
	cnt := 0                     // 회선의 개수

	// 역추적 정보를 사용해 i번 컴퓨터와 연결된 컴퓨터,
	// line{i, minEdge[i]} 또는 line{minEdge[i], i}가 맵에 저장되어 있는지 확인
	// 없다면 맵에 저장하고 회선의 개수를 1늘려준다
	for i := 2; i <= n; i++ {
		if !check[line{i, minEdge[i]}] && !check[line{minEdge[i], i}] {
			check[line{i, minEdge[i]}] = true
			cnt += 1
		}
	}

	fmt.Fprintln(writer, cnt)

	for k := range check {
		fmt.Fprintf(writer, "%d %d\n", k.s, k.e)
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
				minEdge[eg.next] = next // 최솟값이 갱신되면 역추적 정보도 갱신
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
