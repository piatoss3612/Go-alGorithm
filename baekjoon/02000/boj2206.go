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
	Map     [1001][1001]byte
	visited [2][1001][1001]bool // visited[벽을 부술 수 있는 횟수][y 좌표][x 좌표]
	dy      = []int{-1, +0, +1, +0}
	dx      = []int{+0, +1, +0, -1}
)

// (1,1)에서 출발하여 (N,M)으로 이동하는 최단 경로 구하기
// 이동하는 중에 최대 1개의 벽을 부술 수 있다
// 이동할 수 있는 칸은 상하좌우

// 난이도: Gold 3
// 메모리: 56504KB
// 시간: 296ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		line := scanBytes()
		for j := 1; j <= M; j++ {
			Map[i][j] = line[j-1]
		}
	}
}

func Solve() {
	// (1,1)에서 깊이 우선 탐색을 실행한 결과를 출력
	fmt.Fprintln(writer, BFS())
}

type Move struct {
	y, x, cnt, breakable int // y 좌표, x 좌표, 이동한 칸의 개수, 벽을 부술 수 있는 횟수
}

// (1,1)에서 깊이 우선 탐색을 진행하면서 벽을 부술 수 있는 경우에는 벽을 부수고 진행한다
// (N, M)에 도달한 경우, 이동한 칸의 개수는 항상 최소이므로 이동한 칸의 개수 curr.cnt를 반환하고 탐색을 종료한다
// (N, M)에 도달하지 못한 경우, -1을 반환한다
func BFS() int {
	move := []Move{}
	move = append(move, Move{1, 1, 1, 1})
	visited[1][1][1] = true

	// (1,1)에서 시작 칸을 포함하여 벽을 1회 부술 수 있는 상태에서 깊이 우선 탐색 시작

	for len(move) > 0 {
		curr := move[0]
		move = move[1:]

		// (N,M)에 도달한 경우
		if curr.y == N && curr.x == M {
			return curr.cnt
		}

		for i := 0; i < 4; i++ {
			ny, nx := curr.y+dy[i], curr.x+dx[i]
			if valid(ny, nx) {
				// (ny,nx)를 아직 방문하지 않았으며 이동할 수 있는 경우
				if Map[ny][nx] == '0' && !visited[curr.breakable][ny][nx] {
					move = append(move, Move{ny, nx, curr.cnt + 1, curr.breakable})
					visited[curr.breakable][ny][nx] = true
				}

				// (ny,nx)를 아직 방문하지 않았으며 벽이라 이동할 수 없지만, 벽을 부수고 이동할 수 있는 경우
				if Map[ny][nx] == '1' && curr.breakable == 1 && !visited[0][ny][nx] {
					move = append(move, Move{ny, nx, curr.cnt + 1, 0})
					// visited[curr.breakable][ny][nx] = true // 이거 왜 맞았다 처리됩니까??
					visited[0][ny][nx] = true
				}

				// 벽으로 막혀있다고 방문 처리를 해버리면
				// 다른 경로에서 벽을 부수고 진행하는 경우를 배제하게 되므로 주의
				// visited[curr.breakable][ny][nx] = true
			}
		}
	}

	// 최단 거리를 찾지 못한 경우
	return -1
}

func valid(y, x int) bool {
	return y >= 1 && y <= N && x >= 1 && x <= M
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
