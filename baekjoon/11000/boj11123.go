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
	T       int
	H, W    int
	grid    [100][100]byte
	visited [100][100]bool
	dy      = []int{0, 0, 1, -1}
	dx      = []int{1, -1, 0, 0}
)

// 11123번: 양 한마리... 양 두마리...
// https://www.acmicpc.net/problem/11123
// 난이도: 실버 2
// 메모리: 4664 KB
// 시간: 24 ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색, 깊이 우선 탐색
// 비고: 동일한 풀이인데 배열 크기를 동적으로 할당하면 틀림. 왜?
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	// Solve()
}

func Setup() {
	T = scanInt()
	for i := 0; i < T; i++ {
		H, W = scanInt(), scanInt()
		grid = [100][100]byte{}
		visited = [100][100]bool{}
		for j := 0; j < H; j++ {
			b := scanBytes()
			for k := 0; k < W; k++ {
				grid[j][k] = b[k]
			}
		}

		Solve()
	}
}

func Solve() {
	ans := 0

	q := [][2]int{}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == '#' && !visited[i][j] {
				visited[i][j] = true
				q = append(q, [2]int{i, j})

				ans++
			}

			for len(q) > 0 {
				y, x := q[0][0], q[0][1]
				q = q[1:]

				for k := 0; k < 4; k++ {
					ny, nx := y+dy[k], x+dx[k]
					if inRange(ny, nx) && grid[ny][nx] == '#' && !visited[ny][nx] {
						visited[ny][nx] = true
						q = append(q, [2]int{ny, nx})
					}
				}
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func inRange(y, x int) bool {
	return 0 <= y && y < H && 0 <= x && x < W
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
