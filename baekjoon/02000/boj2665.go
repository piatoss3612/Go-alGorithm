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
	board   [51][51]int
	visited [51][51]bool
	dy      = [4]int{0, 0, 1, -1}
	dx      = [4]int{1, -1, 0, 0}
)

// 난이도: Gold 4
// 메모리: 1016KB
// 시간: 4ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색, 0-1 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()

	for i := 1; i <= N; i++ {
		line := scanString()
		for j := 1; j <= N; j++ {
			board[i][j] = int(line[j-1] - '0')
		}
	}
}

// 방의 좌표와 검은 방을 흰 방으로 변경한 횟수를 저장하는 구조체
type Room struct {
	x, y, c int
}

// 검은 방을 흰 방으로 변경한 횟수가 작은 순서대로 정렬되는 우선순위 큐
type PQ []*Room

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].c < pq[j].c
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	item := x.(*Room)
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
	heap.Push(pq, &Room{1, 1, 0})
	visited[1][1] = true

	// 검은 방을 흰 방으로 변경한 횟수가 작은 순서대로 탐색
	for pq.Len() > 0 {
		room := heap.Pop(pq).(*Room)
		x, y, c := room.x, room.y, room.c

		// 목적지에 도착하면 검은 방을 흰 방으로 변경한 횟수를 출력하고 종료
		if x == N && y == N {
			fmt.Fprintln(writer, c)
			return
		}

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx < 1 || nx > N || ny < 1 || ny > N || visited[nx][ny] {
				continue
			}

			visited[nx][ny] = true

			var nc int
			// 흰 방이면 변경 횟수를 그대로, 검은 방이면 변경 횟수를 1 증가
			if board[nx][ny] == 1 {
				nc = c
			} else {
				nc = c + 1
			}

			heap.Push(pq, &Room{nx, ny, nc})
		}
	}
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
