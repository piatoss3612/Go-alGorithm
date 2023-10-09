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
	home    [17][17]int
	// dp[i][j][k]: 파이프의 한쪽 끝이 (i,j)에 있고 파이프의 상태가 k일 때, 파이프의 한쪽 끝이 (N,N)으로 이동시키는 경우의 수
	// 파이프의 상태 0: 가로, 1: 세로, 2: 대각선
	dp   [17][17][3]int
	move = [][]int{{0, 1}, {1, 0}, {1, 1}} // move[0]: 가로로 이동, move[1]: 세로로 이동, move[2]: 대각선으로 이동
)

// 난이도: Gold 5
// 메모리: 17452KB
// 시간: 708ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			home[i][j] = scanInt()
		}
	}
}

func Solve() {
	fmt.Fprintln(writer, rec(1, 2, 0))
}

// y: y좌표
// x: x좌표
// state: 파이프의 상태(0: 가로, 1: 세로, 2: 대각선)
// 파이프의 한쪽 끝이 (y, x)에 있고 상태가 state일 때, 파이프의 한쪽 끝이 (N, N)에 도달하는 경우의 수를 찾는 재귀함수
func rec(y, x, state int) int {
	// 기저 사례: 파이프의 한쪽 끝이 (N, N)으로 이동한 경우
	if y == N && x == N {
		return 1
	}

	ret := &dp[y][x][state]
	if *ret != 0 {
		return *ret
	}

	var next []int // 이동가능한 방향

	switch state {
	case 0: // 파이프가 가로인 상태
		next = append(next, 0, 2)
	case 1: // 파이프가 세로인 상태
		next = append(next, 1, 2)
	case 2: // 파이프가 대각선인 상태
		next = append(next, 0, 1, 2)
	}

	for _, dir := range next {
		ny, nx := y+move[dir][0], x+move[dir][1]

		// 이동할 좌표가 유효한 경우
		if valid(ny, nx) {
			var movable bool
			switch dir {
			case 2:
				movable = moveDiagonal(ny, nx) // 대각선으로 이동할 수 있는지 판별
			default:
				movable = moveColOrRow(ny, nx) // 가로 또는 세로로 이동할 수 있는지 판별
			}

			// 이동할 수 있는 경우
			if movable {
				*ret += rec(ny, nx, dir)
			}
		}
	}

	return *ret
}

// 유효한 좌표인지 판별한다
func valid(y, x int) bool {
	return y >= 1 && y <= N && x >= 1 && x <= N
}

// 가로 또는 세로로 이동할 경우, 이동할 칸이 빈 칸이어야 한다
func moveColOrRow(y, x int) bool {
	return home[y][x] == 0
}

// 대각선으로 이동하려면 이동할 칸과 주변의 두 칸이 유효하고 빈 칸이어야 한다
func moveDiagonal(y, x int) bool {
	return home[y][x] == 0 && valid(y-1, x) && home[y-1][x] == 0 && valid(y, x-1) && home[y][x-1] == 0
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
