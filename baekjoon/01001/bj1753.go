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
	INF     = 9876543210
	edges   [20001][]edge // 각 노드에 연결된 간선의 정보
	dist    [20001]int    // 특정 노드에서 시작해서 다른 모든 노드로 최단 경로
	n, m    int
)

// 간선의 정보: 연결된 노드, 가중치
type edge struct {
	next, weight int
}

// 메모리: 21248KB
// 시간: 156ms
// 우선순위 큐를 사용한 다익스트라 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	s := scanInt()

	var a, b, c int
	for i := 1; i <= m; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()
		edges[a] = append(edges[a], edge{b, c})
	}

	dijkstra(s) // 다익스트라 알고리즘 실행

	// 최단 경로 출력
	for i := 1; i <= n; i++ {
		if dist[i] == INF {
			fmt.Fprintln(writer, "INF")
		} else {
			fmt.Fprintln(writer, dist[i])
		}
	}
}

// heap 패키지를 사용하기 위한 우선 순위 큐 구현
type PQ []edge

func (q PQ) Len() int { return len(q) }

/*
1. Less 함수에서 q[i].weight > q[j].weight를 검사하고 음수 가중치를 사용하는 경우: 최대 힙
2. Less 함수에서 q[i].weight < q[j].weight를 사용하고 양수 가중치를 사용하는 경우: 최소 힙

1의 경우는 메모리: 21248KB, 시간: 156ms
2의 경우는 메모리: 26388KB, 시간: 164ms

입력값에 따른 정렬 시간과 메모리의 차이가 있는 것으로 보인다
*/
func (q PQ) Less(i, j int) bool {
	return q[i].weight > q[j].weight
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

// 노드 v에서 다른 모든 노드로 가는 최단 거리
func dijkstra(v int) {
	// 최솟값을 구하기 위해 길이를 무한으로 초기화
	for i := 1; i <= n; i++ {
		dist[i] = INF
	}
	dist[v] = 0 // 자기 자신으로 가는 최단 거리는 0

	pq := &PQ{{v, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		top := heap.Pop(pq).(edge)
		next := top.next
		cost := -top.weight // 음수로 저장된 가중치를 다시 양수로 변경

		// v에서 next로 가는 가중치가 이미 cost보다 작은 경우
		if dist[next] < cost {
			continue
		}

		// next를 거쳐 다른 노드로 가는 최단 거리를 갱신
		for _, eg := range edges[next] {
			if dist[eg.next] > cost+eg.weight {
				dist[eg.next] = cost + eg.weight
				heap.Push(pq, edge{eg.next, -dist[eg.next]}) // 가중치를 음수로 저장하여 최단 경로를 먼저 꺼내올 수 있도록 한다
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
