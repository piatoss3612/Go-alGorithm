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
	N, E    int
	edges   [801][]Node // i번 정점과 연결된 간선들의 정보
	V1, V2  int         // 1번 정점에서 N번 정점으로 최단 거리를 찾는 중에 반드시 방문해야 하는 두 정점
	dist    []int       // 정점 i에서 시작하여 다른 모든 정점으로 가는 최단 거리
)

const INF = 987654321

type Node struct {
	num, dist int
}

// 난이도: Gold 4
// 메모리: 11432KB
// 시간: 80ms
// 분류: 데이크스트라
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, E = scanInt(), scanInt()
	for i := 1; i <= E; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		// 방향이 없는 간선
		edges[a] = append(edges[a], Node{b, c})
		edges[b] = append(edges[b], Node{a, c})
	}

	V1, V2 = scanInt(), scanInt()
}

func Solve() {
	// 1번 정점에서 N번 정점으로 이동하는 최단 거리를 데이크스트라 알고리즘을 사용하여 찾는다
	// 그런데 중간에 반드시 V1, V2 2개의 정점을 방문해야 한다
	// 이 조건에 맞는 답이 존재하려면 항상 아래 2개의 경우 중 한 가지가 성립되어야 한다

	// # 1. 1 -> V1 -> V2 -> N
	// 1에서 V1으로 가는 최단 거리가 존재한다
	// V1에서 V2로 가는 최단 거리가 존재한다
	// V2에서 N으로 가는 최단 거리가 존재한다
	// 1에서 V1과 V2를 거쳐 N으로 가는 최단 거리 = 1에서 V1으로 가는 최단 거리 + V1에서 V2로 가는 최단 거리 + V2에서 N으로 가는 최단 거리

	// # 2. 1 -> V2 -> V1 -> N
	// 1에서 V2로 가는 최단 거리가 존재한다
	// V2에서 V1으로 가는 최단 거리가 존재한다
	// V1에서 N으로 가는 최단 거리가 존재한다
	// 1에서 V2와 V1을 거쳐 N으로 가는 최단 거리 = 1에서 V2로 가는 최단 거리 + V2에서 V1으로 가는 최단 거리 + V1에서 N으로 가는 최단 거리

	// 1번과 2번이 모두 성립하는 경우는 각각의 값을 비교하여 최단 거리를 구할 수 있다

	// 1번 정점에서 V1, V2로 가는 최단 거리 구하기
	oneToV1, oneToV2 := INF, INF
	Dijkstra(1)
	if dist[V1] != INF {
		oneToV1 = dist[V1]
	}
	if dist[V2] != INF {
		oneToV2 = dist[V2]
	}

	// V1에서 V2로, V2에서 N으로 가는 최단 거리 구하기
	V1ToV2 := INF
	oneV1V2N := INF
	if oneToV1 != INF {
		Dijkstra(V1)
		if dist[V2] != INF {
			V1ToV2 = dist[V2]
			Dijkstra(V2)
			if dist[N] != INF {
				oneV1V2N = oneToV1 + V1ToV2 + dist[N]
			}
		}
	}

	// V2에서 V1으로, V1에서 N으로 가는 최단 거리 구하기
	V2ToV1 := INF
	oneV2V1N := INF
	if oneToV2 != INF {
		Dijkstra(V2)
		if dist[V1] != INF {
			V2ToV1 = dist[V1]
			Dijkstra(V1)
			if dist[N] != INF {
				oneV2V1N = oneToV2 + V2ToV1 + dist[N]
			}
		}
	}

	ans := min(oneV1V2N, oneV2V1N)
	if ans == INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type PQ []*Node

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// s: 데이크스트라 알고리즘으로 최단 거리를 찾기 위한 시작 지점의 번호
// 우선순위 큐를 사용한 데이크스트라 알고리즘을 통해 최단 거리를 찾는다
func Dijkstra(s int) {
	dist = make([]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = INF
	}
	dist[s] = 0

	pq := new(PQ)
	heap.Init(pq)
	heap.Push(pq, &Node{s, 0})

	for len(*pq) > 0 {
		top := heap.Pop(pq).(*Node)
		curr := top.num
		cost := top.dist

		if dist[curr] < cost {
			continue
		}

		for _, next := range edges[curr] {
			if dist[next.num] > next.dist+cost {
				dist[next.num] = next.dist + cost
				heap.Push(pq, &Node{next.num, dist[next.num]})
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
