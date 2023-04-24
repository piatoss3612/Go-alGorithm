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

	N, M, T int
	castle  [101][101]int
	visited [2][101][101]bool // 무기를 들고 있는지 여부에 따라 0, 1

	dy = []int{-1, +0, +1, +0}
	dx = []int{+0, +1, +0, -1}
)

// 난이도: Gold 5
// 메모리: 2260KB
// 시간: 8ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, T = scanInt(), scanInt(), scanInt()

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			castle[i][j] = scanInt()
		}
	}
}

type Hero struct {
	y, x   int
	wasted int
	armed  int
}

const INF = 987654321

func Solve() {
	q := []Hero{}
	q = append(q, Hero{y: 1, x: 1, wasted: 0, armed: 0})
	visited[0][1][1] = true

	minTime := INF

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		y, x := curr.y, curr.x
		armed := curr.armed
		wasted := curr.wasted

		if y == N && x == M {
			minTime = min(minTime, wasted)
			continue
		}

		for i := 0; i < 4; i++ {
			ny, nx := y+dy[i], x+dx[i]

			if valid(ny, nx) && !visited[armed][ny][nx] {
				switch castle[ny][nx] {
				case 0:
					visited[armed][ny][nx] = true
					q = append(q, Hero{y: ny, x: nx, wasted: wasted + 1, armed: armed})
				case 1:
					if armed == 1 {
						visited[armed][ny][nx] = true
						q = append(q, Hero{y: ny, x: nx, wasted: wasted + 1, armed: armed})
					}
				case 2:
					visited[1][ny][nx] = true
					q = append(q, Hero{y: ny, x: nx, wasted: wasted + 1, armed: 1})
				}
			}
		}
	}

	if minTime > T {
		fmt.Fprintln(writer, "Fail")
	} else {
		fmt.Fprintln(writer, minTime)
	}
}

func valid(y, x int) bool {
	return y >= 1 && y <= N && x >= 1 && x <= M
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
