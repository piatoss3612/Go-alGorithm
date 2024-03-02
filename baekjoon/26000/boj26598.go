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
	N, M    int
	graph   [][]byte
	visited [][]bool
	dx      = []int{0, 0, 1, -1}
	dy      = []int{1, -1, 0, 0}
)

// 26598번: 색종이와 공예
// hhttps://www.acmicpc.net/problem/26598
// 난이도: 골드 5
// 메모리: 15104 KB
// 시간: 108 ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	graph = make([][]byte, N)
	for i := 0; i < N; i++ {
		b := scanBytes()
		graph[i] = make([]byte, M)
		for j := 0; j < M; j++ {
			graph[i][j] = b[j]
		}
	}
	visited = make([][]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, M)
	}
}

func Solve() {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if !visited[i][j] {
				if !BFS(i, j) {
					fmt.Fprintln(writer, "BaboBabo")
					return
				}
			}
		}
	}

	fmt.Fprintln(writer, "dd")
}

func BFS(x, y int) bool {
	q := make([][2]int, 0)
	q = append(q, [2]int{x, y})
	visited[x][y] = true
	cnt := 1
	minX, minY, maxX, maxY := x, y, x, y

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]

		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]
			if nx < 0 || nx >= N || ny < 0 || ny >= M {
				continue
			}

			if !visited[nx][ny] && graph[nx][ny] == graph[x][y] {
				visited[nx][ny] = true
				q = append(q, [2]int{nx, ny})

				cnt++

				minX = min(minX, nx)
				minY = min(minY, ny)
				maxX = max(maxX, nx)
				maxY = max(maxY, ny)
			}
		}
	}

	area := (maxX - minX + 1) * (maxY - minY + 1)

	return area == cnt

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
