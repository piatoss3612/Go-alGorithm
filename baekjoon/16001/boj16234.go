package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N, L, R int
	land    [51][51]int
	visited [51][51]bool

	dy = [4]int{0, 0, -1, 1}
	dx = [4]int{-1, 1, 0, 0}
)

// 난이도: Gold 5
// 메모리: 4956KB
// 시간: 120ms
// 분류: 구현, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, L, R = scanInt(), scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			land[i][j] = scanInt()
		}
	}
}

func Solve() {
	cnt := 0

	for {
		visited = [51][51]bool{}

		// 인구 이동이 일어나는지 확인
		moved := false
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if !visited[i][j] {
					if Move(i, j) {
						moved = true
					}
				}
			}
		}

		// 인구 이동이 일어나지 않았다면 종료
		if !moved {
			break
		}

		cnt++
	}

	fmt.Fprintln(writer, cnt)
}

type Point struct {
	y, x int
}

func Move(y, x int) bool {
	q := []Point{{y, x}}
	visited[y][x] = true
	acc := land[y][x]
	union := []Point{{y, x}}

	// BFS
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for i := 0; i < 4; i++ {
			ny, nx := p.y+dy[i], p.x+dx[i]
			if !inRange(ny, nx) || visited[ny][nx] {
				continue
			}

			diff := abs(land[p.y][p.x] - land[ny][nx])
			if L <= diff && diff <= R {
				q = append(q, Point{ny, nx})
				visited[ny][nx] = true
				acc += land[ny][nx]
				union = append(union, Point{ny, nx})
			}
		}
	}

	// 인구 이동이 일어나지 않았다면 종료
	if len(union) == 1 {
		return false
	}

	// 인구 이동
	avg := acc / len(union)

	for _, p := range union {
		land[p.y][p.x] = avg
	}

	return true
}

func inRange(y, x int) bool {
	return 1 <= y && y <= N && 1 <= x && x <= N
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
