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
	N, M    int
	conn    [][]Edge
	S, T    int
)

type Edge struct {
	e, c int
}

type PQ []Edge

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].c < pq[j].c
}
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Edge))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

const INF = 987654321

// 난이도: Gold 5
// 메모리: 9612KB
// 시간: 60ms
// 분류: 데이크스트라
// 풀이: 데이크스트라 알고리즘을 이용하여 S에서 T로 가는 최단 경로를 구한다.
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	conn = make([][]Edge, N+1)
	for i := 1; i <= M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		conn[a] = append(conn[a], Edge{b, c})
		conn[b] = append(conn[b], Edge{a, c})
	}
	S, T = scanInt(), scanInt()
}

func Solve() {
	dist := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = INF
	}
	dist[S] = 0

	pq := new(PQ)
	heap.Init(pq)
	heap.Push(pq, Edge{S, 0})

	for pq.Len() > 0 {
		x := heap.Pop(pq).(Edge)

		if dist[x.e] < x.c {
			continue
		}

		for _, y := range conn[x.e] {
			next := y.e
			cost := y.c

			if dist[next] > dist[x.e]+cost {
				dist[next] = dist[x.e] + cost
				heap.Push(pq, Edge{next, dist[next]})
			}
		}
	}

	fmt.Fprintln(writer, dist[T])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
