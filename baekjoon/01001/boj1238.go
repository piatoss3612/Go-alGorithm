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
	graph   [1001][1001]int // a -> b로 정방향 그래프
	reverse [1001][1001]int // b -> a로 역방향 그래프
	n, m, x int
)

// 메모리: 16596KB
// 시간: 16ms
// 첫 시도는 각각의 정점에서 다른 정점으로 최단거리를 모두 구하고 i->x의 최단거리 + x->i의 최단거리를 합한 값의 최댓값을 구했다
// 답은 맞았지만 시간이 1156ms가 걸렸는데, 다른 분들 풀이가 8, 16ms 정도로 굉장히 짧은 것을 보고 참고하여 다시 풀어보았다
// 첫 시도: https://www.acmicpc.net/source/43366016
// 재시도: https://www.acmicpc.net/source/43366768

// 재시도한 풀이는 정방향 그래프와 역방향 그래프를 사용하여 x에서 출발하는 최단거리를 구해서 더하는 식으로
// 모든 정점에서 다익스트라를 실행하지 않고 2번만 실행하면 된다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m, x = scanInt(), scanInt(), scanInt()
	var from, to, time int
	for i := 1; i <= m; i++ {
		from, to, time = scanInt(), scanInt(), scanInt()
		graph[from][to] = time
		reverse[to][from] = time
	}

	backHome := dijkstra(x, 1) // 정방향 그래프를 사용해서 x번 정점에서 파티가 끝나고 되돌아오는 최단 거리를 구한다
	toParty := dijkstra(x, 2)  // 역방향 그래프를 사용해서 x번을 제외한 각 정점에서 x번 정점으로 가는 최단 거리를 구한다

	max := 0
	for i := 1; i <= n; i++ {
		if i != x && toParty[i]+backHome[i] > max {
			max = toParty[i] + backHome[i]
		}
	}
	fmt.Fprintln(writer, max)
}

type edge struct {
	cost int
	node int
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

func dijkstra(v, way int) []int {
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = INF
	}
	dist[v] = 0

	q := &PQ{{0, v}}
	heap.Init(q)

	var temp *[1001][1001]int

	switch way {
	case 1:
		temp = &graph
	case 2:
		temp = &reverse
	}

	var top edge
	var cost, node int
	for q.Len() > 0 {
		top = heap.Pop(q).(edge)
		cost, node = -top.cost, top.node

		if dist[node] < cost {
			continue
		}

		for i := 1; i <= n; i++ {
			if temp[node][i] > 0 && cost+temp[node][i] < dist[i] {
				dist[i] = cost + temp[node][i]
				heap.Push(q, edge{-dist[i], i})
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
