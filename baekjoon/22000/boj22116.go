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
	N       int
	area    [1001][1001]int
	dx      = []int{0, 0, 1, -1}
	dy      = []int{1, -1, 0, 0}
)

type Item struct {
	x, y, cost int
}

type PQ []*Item

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

const INF = 9876543210

// 22116번: 창영이와 퇴근
// hhttps://www.acmicpc.net/problem/22116
// 난이도: Gold 4
// 메모리: 53688 KB
// 시간: 1200 ms
// 분류: 데이크스트라, 최단 경로
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			area[i][j] = scanInt()
		}
	}
}

func Solve() {
	cost := [1001][1001]int{}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			cost[i][j] = INF
		}
	}

	cost[1][1] = 0

	pq := new(PQ)
	heap.Push(pq, &Item{1, 1, 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)

		if cost[item.x][item.y] < item.cost {
			continue
		}

		for i := 0; i < 4; i++ {
			nx, ny := item.x+dx[i], item.y+dy[i]
			if !InRange(nx, ny) {
				continue
			}

			nCost := max(item.cost, abs(area[item.x][item.y]-area[nx][ny]))
			if cost[nx][ny] > nCost {
				cost[nx][ny] = nCost
				heap.Push(pq, &Item{nx, ny, nCost})
			}
		}
	}

	fmt.Fprintln(writer, cost[N][N])
}

func InRange(x, y int) bool {
	return 1 <= x && x <= N && 1 <= y && y <= N
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
