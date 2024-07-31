package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, x, y, K int
	dx         = [8]int{1, 2, 2, 1, -1, -2, -2, -1}
	dy         = [8]int{2, 1, -1, -2, -2, -1, 1, 2}
)

// 15488번: 나이트가 체스판을 벗어나지 않을 확률
// hhttps://www.acmicpc.net/problem/15488
// 난이도: 골드 5
// 메모리: 960 KB
// 시간: 8 ms
// 분류: 다이나믹 프로그래밍, 확률론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, x, y, K = scanInt(), scanInt(), scanInt(), scanInt()
}

func Solve() {
	// 2개의 2차원 배열을 번갈아가면서 사용
	// dp[k%2][i][j] = k번째 이동 후 (i, j)에 도착할 확률
	dp := make([][][]float64, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]float64, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]float64, N)
		}
	}

	dp[0][x-1][y-1] = 1 // 시작 위치를 확률 1로 초기화
	// K번 이동
	for k := 0; k < K; k++ {
		// 각 이동마다 이전 이동의 확률을 8방향으로 나눠서 다음 이동의 확률을 계산
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				// 현 위치에서 이동할 확률이 0이면 계산할 필요가 없음
				if dp[k%2][i][j] == 0 {
					continue
				}

				// 나이트가 이동할 수 있는 8가지 이동 형태를 탐색
				for d := 0; d < 8; d++ {
					nx, ny := i+dx[d], j+dy[d]
					// 체스판을 벗어나는 경우는 제외
					if nx < 0 || nx >= N || ny < 0 || ny >= N {
						continue
					}

					// 현 위치에서 다음 위치로 이동할 확률을 더해줌
					dp[(k+1)%2][nx][ny] += dp[k%2][i][j] / 8
				}
			}
		}

		// 다음 이동을 위해 현재 체스판을 초기화
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				dp[k%2][i][j] = 0
			}
		}
	}

	// K번 이동 후 모든 위치의 확률을 더함
	ans := 0.0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ans += dp[K%2][i][j]
		}
	}

	fmt.Fprintln(writer, ans)
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
