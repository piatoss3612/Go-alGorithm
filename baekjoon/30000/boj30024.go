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
	N, M, K int
	field   [1001][1001]int
	visited [1001][1001]bool
	corns   *Corns

	dy = [4]int{0, 0, 1, -1}
	dx = [4]int{1, -1, 0, 0}
)

// y, x 좌표를 가지는 Corn
type Corn struct {
	y, x int
}

type Corns []*Corn // Corn의 slice, 우선순위 큐로 사용

func (c Corns) Len() int {
	return len(c)
}
func (c Corns) Less(i, j int) bool {
	return field[c[i].y][c[i].x] > field[c[j].y][c[j].x] // field 값이 큰 Corn이 우선순위가 높음
}
func (c Corns) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c *Corns) Push(x interface{}) {
	*c = append(*c, x.(*Corn))
}
func (c *Corns) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	corns = new(Corns)
	heap.Init(corns)

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			field[i][j] = scanInt()

			if i == 1 || i == N || j == 1 || j == M { // 가장자리에 있는 Corn은 미리 heap에 넣어둠
				heap.Push(corns, &Corn{i, j})
				visited[i][j] = true // 방문 처리
			}
		}
	}

	K = scanInt()
}

func Solve() {
	for i := 0; i < K; i++ {
		corn := heap.Pop(corns).(*Corn)
		fmt.Fprintln(writer, corn.y, corn.x) // 수확할 수 있는 Corn 중 가장 field 값이 큰 Corn을 출력

		// 수확한 Corn의 상하좌우 Corn을 heap에 넣음
		for j := 0; j < 4; j++ {
			ny, nx := corn.y+dy[j], corn.x+dx[j]
			if inRange(ny, nx) && !visited[ny][nx] {
				heap.Push(corns, &Corn{ny, nx})
				visited[ny][nx] = true
			}
		}

	}
}

func inRange(y, x int) bool {
	return 1 <= y && y <= N && 1 <= x && x <= M
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
