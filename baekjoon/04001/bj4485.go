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
	graph   [126][126]int
	dist    [126][126]int
	dx      = []int{0, 0, 1, -1}
	dy      = []int{1, -1, 0, 0}
	n       int
	cnt     int
)

type edge struct {
	cost int
	x, y int
}

// 녹색 옷 입은 애는 링크입니다...

// 메모리: 2960KB
// 시간: 16ms
// 다익스트라 알고리즘
// 최단 거리를 찾기 위해 지나갔던 길을 다시 돌아갈 수도 있으므로
// BFS로 방문 여부를 확인하면서 진행하는 방법은 사용할 수 없다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for {
		n = scanInt()
		if n == 0 {
			return
		}
		cnt += 1

		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				graph[i][j] = scanInt()
				dist[i][j] = INF
			}
		}

		dijkstra()

		fmt.Fprintf(writer, "Problem %d: %d\n", cnt, dist[n][n])
	}
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

func dijkstra() {
	dist[1][1] = graph[1][1]

	q := &PQ{{-graph[1][1], 1, 1}}
	heap.Init(q)

	var top edge
	var cost, x, y int
	var nx, ny int
	for q.Len() > 0 {
		top = heap.Pop(q).(edge)
		cost = -top.cost
		x, y = top.x, top.y

		if dist[x][y] < cost {
			continue
		}

		// 상하좌우 경로 검사
		for i := 0; i < 4; i++ {
			nx, ny = x+dx[i], y+dy[i]
			if isValid(nx, ny) {
				if cost+graph[nx][ny] < dist[nx][ny] {
					dist[nx][ny] = cost + graph[nx][ny]
					heap.Push(q, edge{-dist[nx][ny], nx, ny})
				}
			}
		}
	}
}

func isValid(x, y int) bool {
	if x >= 1 && x <= n && y >= 1 && y <= n {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
