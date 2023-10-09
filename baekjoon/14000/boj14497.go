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

	N, M           int
	x1, y1, x2, y2 int
	class          [301][301]byte
	visited        [301][301]bool

	dy = []int{-1, +0, +1, +0}
	dx = []int{+0, +1, +0, -1}
)

// 난이도: Gold 4
// 메모리: 6192KB
// 시간: 40ms
// 분류: 너비 우선 탐색, 우선순위 큐
// 시간 복잡도: O(NM)
// 공간 복잡도: O(NM)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	x1, y1, x2, y2 = scanInt(), scanInt(), scanInt(), scanInt()

	for i := 1; i <= N; i++ {
		line := scanBytes()
		for j, b := range line {
			class[i][j+1] = b
		}
	}
}

type Jump struct {
	x, y int
	try  int
}

type PQ []Jump

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].try < pq[j].try // 점프 횟수가 작은 경로를 우선적으로 탐색한다.
}
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Jump))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func Solve() {
	/*
		# 풀이 과정

		1. 주난이가 점프한다.
		2. 상하좌우로 파동이 퍼져나가 친구들을 쓰러뜨린다. 친구를 쓰러트리면 파동이 멈춘다.
		3. 초코바를 찾을 때까지 1, 2를 반복한다.

		4. 1, 2를 반복하는 대신, 친구가 쓰러지면 그 지점에서 점프 횟수를 1 늘리고 파동이 멈추지 않고 계속 퍼져나가도록 한다.
		5. 점프 횟수가 작은 이동 경로를 우선적으로 탐색할 수 있도록 우선순위 큐를 사용한다.
	*/

	pq := new(PQ)
	heap.Init(pq)
	heap.Push(pq, Jump{x: x1, y: y1, try: 1})
	visited[x1][y1] = true

	for pq.Len() > 0 {
		j := heap.Pop(pq).(Jump)

		x, y := j.x, j.y
		try := j.try

		// 초코바를 찾은 경우 *항상 초코바를 찾을 수 있다*
		if x == x2 && y == y2 {
			fmt.Fprintln(writer, try)
			return
		}

		// 상하좌우로 파동이 퍼져나간다.
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if valid(nx, ny) && !visited[nx][ny] {
				if class[nx][ny] == '1' {
					heap.Push(pq, Jump{x: nx, y: ny, try: try + 1})
				} else {
					heap.Push(pq, Jump{x: nx, y: ny, try: try})
				}
				visited[nx][ny] = true
			}
		}
	}
}

func valid(x, y int) bool {
	return x >= 1 && x <= N && y >= 1 && y <= M
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
