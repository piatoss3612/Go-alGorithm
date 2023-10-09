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
	V, E, P int
	edges   [][]Edge
	dist    []int
	visited []bool
)

type Edge struct {
	n, c int
}

const INF = 987654321

// 난이도: Gold 4
// 메모리: 2344KB
// 시간: 12ms
// 분류: 데이크스트라
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	V, E, P = scanInt(), scanInt(), scanInt()
	edges = make([][]Edge, V+1)
	dist = make([]int, V+1)
	visited = make([]bool, V+1)
	for i := 1; i <= E; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		edges[a] = append(edges[a], Edge{b, c})
		edges[b] = append(edges[b], Edge{a, c})
	}
}

func Solve() {
	// 데이크스트라 알고리즘으로 1에서 V까지의 최단거리 구하기
	Dijkstra()

	// 1에서 V로 가는 경로 사이에 P를 방문하는 경로가 있는지 역추적
	ans := Track(V)
	if ans {
		fmt.Fprintln(writer, "SAVE HIM")
	} else {
		fmt.Fprintln(writer, "GOOD BYE")
	}

	/*
		# 다른 풀이
		// 메모리: 2800KB
		// 시간: 16ms

		(1 -> V 경로의 최단거리)와 (1 -> P 경로의 최단거리 + P -> V 경로의 최단거리)를 비교

		func Solve() {
			a := Dijkstra(1, V)
			b := Dijkstra(1, P)
			c := Dijkstra(P, V)

			if a == b+c {
				fmt.Fprintln(writer, "SAVE HIM")
			} else {
				fmt.Fprintln(writer, "GOOD BYE")
			}
		}

		func Dijkstra(from, to int) int {
			dist := make([]int, V+1)
			for i := 1; i <= V; i++ {
				dist[i] = INF
			}
			dist[from] = 0

			q := new(PQ)
			heap.Init(q)
			heap.Push(q, &Edge{from, 0})

			for q.Len() > 0 {
				x := heap.Pop(q).(*Edge)

				here := x.n
				acc := x.c

				if dist[here] < acc {
					continue
				}

				for _, connected := range edges[here] {
					next := connected.n
					extra := connected.c

					if dist[next] > acc+extra {
						dist[next] = acc + extra
						heap.Push(q, &Edge{next, dist[next]})
					}
				}
			}

			return dist[to]
		}
	*/
}

type PQ []*Edge

func (q PQ) Len() int { return len(q) }
func (q PQ) Less(i, j int) bool {
	return q[i].c < q[j].c
}
func (q PQ) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *PQ) Push(x interface{}) {
	*q = append(*q, x.(*Edge))
}
func (q *PQ) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func Dijkstra() {
	for i := 1; i <= V; i++ {
		dist[i] = INF
	}
	dist[1] = 0 // 시작점

	q := new(PQ)
	heap.Init(q)
	heap.Push(q, &Edge{1, 0})

	for q.Len() > 0 {
		x := heap.Pop(q).(*Edge)

		here := x.n
		acc := x.c

		// 최단경로보다 더 길다면 무시
		if dist[here] < acc {
			continue
		}

		// 인접한 정점들을 검사
		for _, connected := range edges[here] {
			next := connected.n
			extra := connected.c

			// 최단경로 갱신
			if dist[next] > acc+extra {
				dist[next] = acc + extra
				heap.Push(q, &Edge{next, dist[next]})
			}
		}
	}
}

func Track(x int) bool {
	if x == P {
		return true
	}

	visited[x] = true
	ans := false

	for _, next := range edges[x] {
		if !visited[next.n] {
			if dist[x]-next.c == dist[next.n] {
				ans = ans || Track(next.n)
			}
		}
	}

	visited[x] = false
	return ans
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
