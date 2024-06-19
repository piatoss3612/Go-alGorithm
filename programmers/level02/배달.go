package level02

import "container/heap"

type Edge struct {
	to   int
	dist int
}

type Node struct {
	pos  int
	dist int
}

type PQ []*Node

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/12978
// 분류: 다익스트라, 최단거리
func solution(N int, road [][]int, k int) int {
	edges := make([][]Edge, N+1)
	for _, r := range road {
		u, v, w := r[0], r[1], r[2]
		edges[u] = append(edges[u], Edge{v, w})
		edges[v] = append(edges[v], Edge{u, w})
	}

	dist := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = 1e9
	}
	dist[1] = 0

	pq := new(PQ)
	heap.Push(pq, &Node{1, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Node)
		for _, e := range edges[cur.pos] {
			if dist[e.to] > cur.dist+e.dist {
				dist[e.to] = cur.dist + e.dist
				heap.Push(pq, &Node{e.to, dist[e.to]})
			}
		}
	}

	cnt := 0
	for i := 1; i <= N; i++ {
		if dist[i] <= k {
			cnt++
		}
	}

	return cnt
}
