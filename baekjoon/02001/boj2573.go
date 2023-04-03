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
	arctic  [301][301]int
	visited [301][301]bool
	queue   []Iceberg
	dy      = []int{-1, +0, +1, +0}
	dx      = []int{+0, +1, +0, -1}
)

type Iceberg struct {
	y, x, h int
}

// 난이도: Gold 4
// 메모리: 12984KB
// 시간: 188ms
// 분류:  구현, 너비 우선 탐색, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			arctic[i][j] = scanInt()
			if arctic[i][j] != 0 {
				queue = append(queue, Iceberg{i, j, arctic[i][j]})
			}
		}
	}
}

func Solve() {
	/*
		1. 깊이 우선 탐색을 통해 빙산이 두 덩어리 이상으로 분리되어 있는지 확인
			1-1. 두 덩어리 이상으로 분리되어 있다면, 빙산이 분리되기까지 걸린 시간을 출력하고 종료
		2. 큐에 저장된 빙산의 정보를 순회하며, 1년간 빙산이 녹는 과정을 시뮬레이션
			2-1. 빙산이 녹는 과정에서 갱신된 정보는 candidates에 저장
			2-2. candidates에 저장된 정보를 바탕으로 arctic 배열을 갱신
		3. 큐가 빌 때까지 1~2 과정을 반복
		4. 모든 빙하가 녹을 때까지 빙산이 분리되지 않았다면, 0을 출력
	*/

	year := 0

	for len(queue) > 0 {
		visited = [301][301]bool{}
		dfsCnt := 0

		for i := 1; i <= N; i++ {
			for j := 1; j <= M; j++ {
				if !visited[i][j] && arctic[i][j] > 0 {
					if dfsCnt > 0 {
						fmt.Fprintln(writer, year)
						return
					} else {
						DFS(i, j)
						dfsCnt++
					}
				}
			}
		}

		candidates := []Iceberg{}

		for len(queue) > 0 {
			front := queue[0]
			queue = queue[1:]

			zeros := 0
			for i := 0; i < 4; i++ {
				ny, nx := front.y+dy[i], front.x+dx[i]
				if valid(ny, nx) && arctic[ny][nx] == 0 {
					zeros++
				}
			}

			front.h -= zeros

			candidates = append(candidates, front)
		}

		for _, c := range candidates {
			if c.h <= 0 {
				arctic[c.y][c.x] = 0
			} else {
				arctic[c.y][c.x] = c.h
				queue = append(queue, c)
			}
		}

		year++
	}

	fmt.Fprintln(writer, 0)
}

func DFS(y, x int) {
	visited[y][x] = true

	for i := 0; i < 4; i++ {
		ny, nx := y+dy[i], x+dx[i]
		if valid(ny, nx) && !visited[ny][nx] && arctic[ny][nx] > 0 {
			DFS(ny, nx)
		}
	}
}

func valid(y, x int) bool {
	return y >= 1 && y <= N && x >= 1 && x <= M
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
