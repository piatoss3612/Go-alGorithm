package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N, M, X, Y int
	roads      [][]Road
)

type Road struct {
	city int
	dist int
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

const INF = 987654321

// 난이도: Gold 4
// 메모리: 7024KB
// 시간: 40ms
// 분류: 데이크스트라
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, X, Y = scanInt(), scanInt(), scanInt(), scanInt()
	roads = make([][]Road, N)

	for i := 1; i <= M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		roads[a] = append(roads[a], Road{city: b, dist: c})
		roads[b] = append(roads[b], Road{city: a, dist: c})
	}
}

func Solve() {
	dist := dijkstra() // Y에서 각 도시까지의 최단거리
	sort.Ints(dist)

	// 떡은 한번에 하나씩만 들고 갈 수 있다.
	// 하루에 X보다 먼 거리를 걷지 않고 거리가 가까운 집부터 방문한다.
	// 잠은 꼭 본인 집에서 자야 하므로 왕복할 수 없는 거리는 다음날 가기로 다짐한다.
	days := 1
	left := X
	i := 1
	for i < N {
		move := dist[i] * 2

		if move > left {
			if left == X {
				fmt.Fprintln(writer, -1)
				return
			}
			days += 1
			left = X
			continue
		}

		left -= move
		i += 1
	}

	fmt.Fprintln(writer, days)
}

func dijkstra() []int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = INF
	}
	dist[Y] = 0

	pq := new(PQ)
	heap.Push(pq, Road{city: Y, dist: 0})

	for pq.Len() > 0 {
		x := heap.Pop(pq).(Road)
		c, d := x.city, x.dist

		if d > dist[c] {
			continue
		}

		for _, next := range roads[c] {
			nc, ex := next.city, next.dist
			if dist[nc] > dist[c]+ex {
				dist[nc] = dist[c] + ex
				heap.Push(pq, Road{city: nc, dist: dist[nc]})
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
