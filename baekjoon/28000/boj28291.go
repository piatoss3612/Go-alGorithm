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
	W, H    int
	N       int
	board   [50][50]int
	dx      = [4]int{0, 0, 1, -1}
	dy      = [4]int{1, -1, 0, 0}
	queue   [][2]int
)

// 28291번: 레드스톤
// hhttps://www.acmicpc.net/problem/28291
// 난이도: 골드 5
// 메모리: 1196 KB
// 시간: 4 ms
// 분류: 너비 우선 탐색, 그래프 이론, 그래프 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	W, H = scanInt(), scanInt()
	N = scanInt()

	for i := 0; i < N; i++ {
		b, x, y := scanString(), scanInt(), scanInt()
		switch b {
		case "redstone_block":
			board[x][y] = 1
			queue = append(queue, [2]int{x, y})
		case "redstone_dust":
			board[x][y] = 2
		case "redstone_lamp":
			board[x][y] = 3
		}
	}
}

const INF = 987654321

func Solve() {
	var visited [50][50]int

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		switch board[x][y] {
		case 1:
			visited[x][y] = 15

			for i := 0; i < 4; i++ {
				nx, ny := x+dx[i], y+dy[i]
				if isInRange(nx, ny) && visited[nx][ny] < 15 {
					visited[nx][ny] = 15
					queue = append(queue, [2]int{nx, ny})
				}
			}
		case 2:
			if visited[x][y] == 1 {
				continue
			}

			for i := 0; i < 4; i++ {
				nx, ny := x+dx[i], y+dy[i]
				if isInRange(nx, ny) && visited[nx][ny] < visited[x][y]-1 {
					visited[nx][ny] = visited[x][y] - 1
					queue = append(queue, [2]int{nx, ny})
				}
			}
		case 3:
			visited[x][y] = INF
		}
	}

	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			if board[i][j] == 3 && visited[i][j] != INF {
				fmt.Fprintln(writer, "failed")
				return
			}
		}
	}

	fmt.Fprintln(writer, "success")
}

func isInRange(x, y int) bool {
	return 0 <= x && x < W && 0 <= y && y < H
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
