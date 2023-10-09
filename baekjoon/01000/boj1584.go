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
	board   [501][501]int
	visited [501][501]bool
	dy      = [4]int{0, 0, 1, -1}
	dx      = [4]int{1, -1, 0, 0}
)

// 난이도: Gold 5
// 메모리: 13624KB
// 시간: 68ms
// 분류: 그래프 이론, 그래프 탐색, 0-1 너비 우선 탐색
// 참고: https://ebabby.tistory.com/5 (Dequeue를 사용하여 풀고자 했으나 시간초과가 발생하여 우선순위 큐를 사용함. Go에서는 Dequeue를 자체적으로 지원하지 않음.)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()

	for k := 1; k <= N; k++ {
		x1, y1, x2, y2 := scanInt(), scanInt(), scanInt(), scanInt()
		x1, x2 = swap(x1, x2)
		y1, y2 = swap(y1, y2)

		// 위험한 구역은 1로 표시
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				board[i][j] = 1
			}
		}
	}

	M = scanInt()

	for k := 1; k <= M; k++ {
		x1, y1, x2, y2 := scanInt(), scanInt(), scanInt(), scanInt()
		x1, x2 = swap(x1, x2)
		y1, y2 = swap(y1, y2)

		// 죽음의 구역은 방문 처리
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				visited[i][j] = true
			}
		}
	}
}

type Point struct {
	x, y, h int
}

// 우선순위 큐 정의
type PQ []*Point

func (pq PQ) Len() int {
	return len(pq)
}

// 필요한 생명력이 적은 위치부터 탐색하기 위해 h를 기준으로 정렬
func (pq PQ) Less(i, j int) bool {
	return pq[i].h < pq[j].h
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	item := x.(*Point)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Solve() {
	pq := new(PQ)
	heap.Init(pq)
	heap.Push(pq, &Point{0, 0, 0})
	visited[0][0] = true

	for pq.Len() != 0 {
		p := heap.Pop(pq).(*Point)
		x, y, h := p.x, p.y, p.h

		// 목적지에 도착하면 필요 생명력 출력 후 종료
		if x == 500 && y == 500 {
			fmt.Fprintln(writer, h)
			return
		}

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]

			if nx < 0 || nx > 500 || ny < 0 || ny > 500 {
				continue
			}

			if visited[nx][ny] {
				continue
			}

			visited[nx][ny] = true

			// 위험한 구역이면 생명력이 1 더 필요함
			var nh int
			if board[nx][ny] == 1 {
				nh = h + 1
			} else {
				nh = h
			}

			heap.Push(pq, &Point{nx, ny, nh}) // 우선순위 큐에 삽입
		}
	}

	fmt.Fprintln(writer, -1) // 목적지에 도착하지 못하면 -1 출력
}

func swap(a, b int) (int, int) {
	if b > a {
		return a, b
	}
	return b, a
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
