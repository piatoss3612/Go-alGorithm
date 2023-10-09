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
	A, B, C int
	roads   [][]Road
)

type Road struct {
	node, dist int
}

const INF = 987654321

// 난이도: Gold 4
// 메모리: 73124KB
// 시간: 732ms
// 분류: 데이크스트라
// 시간복잡도: O(NlogN)
// 공간복잡도: O(M^2)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	roads = make([][]Road, N+1)
	A, B, C = scanInt(), scanInt(), scanInt()

	M = scanInt()
	for i := 1; i <= M; i++ {
		d, e, l := scanInt(), scanInt(), scanInt()
		roads[d] = append(roads[d], Road{e, l})
		roads[e] = append(roads[e], Road{d, l})
	}
}

func Solve() {
	// A, B, C를 각각 시작점으로 두고 데이크스트라 알고리즘을 수행한다.
	a := Dijkstra(A)
	b := Dijkstra(B)
	c := Dijkstra(C)

	// 각 노드에서 A, B, C 까지의 거리 중 최솟값이 가장 큰 노드를 찾는다.
	ans := -1
	maxDist := 0
	for i := 1; i <= N; i++ {
		dist := min(min(a[i], b[i]), c[i])
		if dist > maxDist {
			ans = i
			maxDist = dist
		}
	}

	fmt.Fprintln(writer, ans)
}

type PQ []Road

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Road))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func Dijkstra(v int) []int {
	pq := new(PQ)
	dist := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = INF
	}
	dist[v] = 0
	heap.Push(pq, Road{v, 0})

	for pq.Len() > 0 {
		x := heap.Pop(pq).(Road)

		if x.dist > dist[x.node] {
			continue
		}

		for _, road := range roads[x.node] {
			if dist[road.node] > x.dist+road.dist {
				dist[road.node] = x.dist + road.dist
				heap.Push(pq, Road{road.node, dist[road.node]})
			}
		}
	}

	return dist
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func parseInt(s string) int {

	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
