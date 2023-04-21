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

	N, M  int
	roads [][]Node
)

type Node struct {
	num  int
	cost int
}

type PQ []Node

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
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
// 메모리: 8616KB
// 시간: 80ms
// 분류: 데이크스트라
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()

	roads = make([][]Node, N+1)

	for i := 1; i <= M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		roads[a] = append(roads[a], Node{num: b, cost: c})
		roads[b] = append(roads[b], Node{num: a, cost: c})
	}
}

func Solve() {
	dist := make([]int, N+1)
	for i := 2; i <= N; i++ {
		dist[i] = INF
	}

	pq := new(PQ)
	heap.Push(pq, Node{num: 1, cost: 0})

	for pq.Len() > 0 {
		n := heap.Pop(pq).(Node)
		from := n.num
		cost := n.cost

		if dist[from] < cost {
			continue
		}

		for _, next := range roads[from] {
			to := next.num
			extra := next.cost

			if dist[to] > cost+extra {
				dist[to] = cost + extra
				heap.Push(pq, Node{num: to, cost: dist[to]})
			}
		}
	}

	fmt.Fprintln(writer, dist[N])
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
