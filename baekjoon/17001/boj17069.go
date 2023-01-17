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
	home    [33][33]int
	dp      [33][33][3]int                    // dp[i][j][k]: 파이프의 끝이 (1,2)에서 가로 상태로 이동을 시작해서 파이프의 끝이 (i, j)에 도달하였으며 파이프의 상태가 k인 경우의 수
	move    = [][]int{{0, 1}, {1, 0}, {1, 1}} // move[0]: 가로로 이동시 좌표 변화, move[1]: 세로로 이동시 좌표 변화, move[2]: 대각선으로 이동시 좌표 변화
)

// 난이도: Gold 4
// 메모리: 1008KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
// 17070번 문제를 재귀함수형 풀이에서 반복문형 풀이로 변경해 풀어보았다
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
	dp[1][2][0] = 1
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			// (i, j) 좌표에서 파이프의 상태가 가로, 세로, 대각선인 경우에 대해 각각 탐색
			for k := 0; k <= 2; k++ {
				if dp[i][j][k] != 0 {
					var next []int

					// 파이프의 상태에 따라 회전시키는 방향 분기 처리
					switch k {
					case 0:
						next = append(next, 0, 2) // 가로인 경우: 가로, 대각선 회전 가능
					case 1:
						next = append(next, 1, 2) // 세로인 경우: 세로, 대각선 회전 가능
					case 2:
						next = append(next, 0, 1, 2) // 대각선인 경우: 가로, 세로, 대각선 회전 가능
					}

					// 파이프 회전 및 이동
					for _, dir := range next {
						ny, nx := i+move[dir][0], j+move[dir][1]

						// 좌표가 범위를 벗어난 경우
						if !valid(ny, nx) {
							continue
						}

						var movable bool

						// 회전 및 이동시킬 방향에 따라 이동가능 여부 확인
						switch dir {
						case 2:
							movable = moveDiagonal(ny, nx) // 대각선으로 회전 및 이동시키기 위해 주변의 칸들이 비어있는지 확인
						default:
							movable = moveColOrRow(ny, nx) // 가로 또는 세로로 회전 및 이동시키기 위해 이동할 칸이 비어있는지 확인
						}

						// 이동할 수 있는 경우
						// 이동하려는 위치에 도달할 수 있는 경우의 수를 현재 위치에 도달할 수 있는 경우의 수만큼 더한다
						if movable {
							dp[ny][nx][dir] += dp[i][j][k]
						}
					}
				}
			}
		}
	}

	fmt.Fprintln(writer, dp[N][N][0]+dp[N][N][1]+dp[N][N][2]) // (N, N)에 도달할 수 있는 모든 경우의 수의 합
}

// (y, x)가 범위 내에 있는 좌표인지 확인
func valid(y, x int) bool {
	return y >= 1 && y <= N && x >= 1 && x <= N
}

// 가로 또는 세로로 이동이 가능한지 확인
func moveColOrRow(y, x int) bool {
	return home[y][x] == 0
}

// 대각선으로 이동이 가능한지 확인
func moveDiagonal(y, x int) bool {
	return home[y][x] == 0 && valid(y-1, x) && home[y-1][x] == 0 && valid(y, x-1) && home[y][x-1] == 0
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
