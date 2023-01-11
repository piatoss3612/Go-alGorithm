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
	L, R, B, K int
	dp         [1000001][4]int // dp[i][j]: i만큼 에너지를 소비한 상태에서 방향 j를 바라보고 있을 때, 제식 수행 횟수의 최솟값
)

const INF = 987654321

// 난이도: Gold 4
// 메모리: 32160KB
// 시간: 64ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	L, R, B, K = scanInt(), scanInt(), scanInt(), scanInt()
}

func Solve() {
	// 최솟값 비교를 위해 dp를 INF로 초기화
	for i := 0; i <= K; i++ {
		dp[i][0] = INF
		dp[i][1] = INF
		dp[i][2] = INF
		dp[i][3] = INF
	}

	// 임의로 0을 시작 방향으로 지정
	// 1: 오른쪽으로 90도
	// 2: 오른쪽으로 180도
	// 3: 오른쪽으로 270도

	dp[0][0] = 0 // 제식 훈련 시작 시의 상태

	for i := 0; i <= K; i++ {
		for j := 0; j <= 3; j++ {
			// i만큼 에너지를 소비했고 방향 j를 바라보고 있는 경우
			// 현재 상태에서 3가지 제식을 각각 실행하는 경우의 상태값 갱신
			if dp[i][j] != INF {
				// 좌로 돌아
				if i+L <= K {
					dir := (j + 3) % 4
					dp[i+L][dir] = min(dp[i+L][dir], dp[i][j]+1)
				}

				// 우로 돌아
				if i+R <= K {
					dir := (j + 1) % 4
					dp[i+R][dir] = min(dp[i+R][dir], dp[i][j]+1)
				}

				// 뒤로 돌아
				if i+B <= K {
					dir := (j + 2) % 4
					dp[i+B][dir] = min(dp[i+B][dir], dp[i][j]+1)
				}
			}
		}
	}

	if dp[K][0] != INF {
		fmt.Fprintln(writer, dp[K][0]) // 정확히 K만큼의 에너지를 소모하고 훈련을 시작한 방향을 바라보고 있는 경우
	} else {
		fmt.Fprintln(writer, -1)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
