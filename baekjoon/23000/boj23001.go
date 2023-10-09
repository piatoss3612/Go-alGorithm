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
	T, R, C int
	house   [301][301]int // 토끼집
	cells   *Cells
	dy      = []int{-1, +0, +1, +0}
	dx      = []int{+0, +1, +0, -1}
)

// 토끼집의 각 좌표
type Cell struct {
	y, x  int // 좌표
	boxes int // 박스의 개수
}

// 우선순위 큐 정의 및 구현
type Cells []*Cell

func (c Cells) Len() int { return len(c) }
func (c Cells) Less(i, j int) bool {
	return c[i].boxes > c[j].boxes
}
func (c Cells) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c *Cells) Push(x interface{}) {
	*c = append(*c, x.(*Cell))
}
func (c *Cells) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

// 난이도: Gold 3
// 메모리: 78172KB
// 시간: 2744ms
// 분류: 우선순위 큐, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Input()
		Solve(i)
	}
}

func Input() {
	R, C = scanInt(), scanInt()
	cells = new(Cells)
	heap.Init(cells)
	for i := 1; i <= R; i++ {
		for j := 1; j <= C; j++ {
			house[i][j] = scanInt()
			heap.Push(cells, &Cell{i, j, house[i][j]})
		}
	}
}

func Solve(caseNo int) {
	cnt := 0

	// 그리디 알고리즘:
	// 박스의 개수가 가장 많은 cellX부터 시작하여
	// cellX와 인접한 셀의 박스의 개수를 cellX의 박스의 개수-1로 조정하는 과정을 반복함으로써
	// 토끼집에 추가할 박스의 개수를 최소화할 수 있다

	for len(*cells) > 0 {
		cell := heap.Pop(cells).(*Cell) // 우선순위 큐에서 박스의 개수가 가많 많은 셀 꺼내오기

		// 실제 박스의 개수와 일치하지 않는 경우
		// 해당 셀은 갱신된 것이므로 현재 정보를 파기한다
		if house[cell.y][cell.x] != cell.boxes {
			continue
		}

		// cell과 인접한 셀 탐색
		for i := 0; i < 4; i++ {
			ny, nx := cell.y+dy[i], cell.x+dx[i]
			// 유효한 좌표인 경우에만
			if valid(ny, nx) {
				// 인접한 셀의 박스의 개수와 cell의 박스의 개수 차이가 1보다 큰 경우
				// 부족한 박스의 개수를 보충하고 갱신된 값을 우선순위 큐에 추가
				if house[ny][nx] < cell.boxes-1 {
					cnt += cell.boxes - 1 - house[ny][nx]
					house[ny][nx] = cell.boxes - 1
					heap.Push(cells, &Cell{ny, nx, house[ny][nx]})
				}
			}
		}
	}

	fmt.Fprintf(writer, "Case #%d: %d\n", caseNo, cnt)
}

func valid(y, x int) bool {
	if y >= 1 && y <= R && x >= 1 && x <= C {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
