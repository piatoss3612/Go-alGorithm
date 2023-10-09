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

	N       int
	board   [2001][2001]byte
	visited [2001][2001]bool
	dy      = [7]int{0, 0, -1, 1, 1, -1, -1}
	dx      = [7]int{1, -1, 0, 1, -1, 1, -1}
	sy, sx  int
)

// 난이도: Silver 1
// 메모리: 23912KB
// 시간: 364ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		b := scanString()
		for j := 1; j <= N; j++ {
			board[i][j] = b[j-1]
			// 도착지점 찾기
			if board[i][j] == 'F' {
				sy, sx = i, j
			}
		}
	}
}

func Solve() {
	// 도착지점으로 갈 수 있는 모든 지점을 찾기 위해
	// 역으로 도착지점에서 출발하여 도달할 수 있는 모든 지점을 찾는다.
	// 이 때, 역으로 움직여야 하므로 위로 움직일 수 없는 제약사항이 아래로 움직일 수 없는 제약사항이 된다.
	// 나머지 이동은 제약사항이 없으므로 7방향으로 움직이며 너비우선탐색을 진행한다.
	cnt := 0

	q := [][2]int{{sy, sx}}
	visited[sy][sx] = true

	for len(q) > 0 {
		y, x := q[0][0], q[0][1]
		q = q[1:]

		for i := 0; i < 7; i++ {
			ny, nx := y+dy[i], x+dx[i]

			if ny < 1 || ny > N || nx < 1 || nx > N {
				continue
			}

			if visited[ny][nx] {
				continue
			}

			if board[ny][nx] == '.' {
				q = append(q, [2]int{ny, nx})
				visited[ny][nx] = true
				cnt++
			}
		}
	}

	fmt.Fprintln(writer, cnt)
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
