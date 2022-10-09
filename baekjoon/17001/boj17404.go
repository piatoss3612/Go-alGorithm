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
	cost    [1001][4]int    // cost[i][j]: i번 집을 j(빨강 1, 초록 2, 파랑 3)색으로 칠하는 비용
	dp      [4][1001][4]int // dp[i][j][k]: 1번 집을 i 색으로 칠하기 시작하여 j번 집을 k 색으로 칠할 때, j~N번 집을 칠하는 비용의 최솟값
)

const INF = 987654321 // 최댓값 비교를 위한 상수

// 메모리: 1340KB
// 시간: 4ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= 3; j++ {
			cost[i][j] = scanInt()
		}
	}

	ans := INF
	// 1번 집을 각각의 색으로 칠하기 시작한 경우의 최솟값 비교
	for i := 1; i <= 3; i++ {
		ans = min(ans, solve(1, i, i)+cost[1][i])
	}

	fmt.Fprintln(writer, ans)
}

func solve(house, currentColor, startColor int) int {
	// 기저 사례: N번째 집을 칠하는 경우
	// N번째 집은 N-1번째 집과 1번째 집의 색과 다른 색으로 칠해야 한다
	if house == N-1 {
		temp := INF
		for i := 1; i <= 3; i++ {
			if currentColor != i && startColor != i {
				temp = min(temp, cost[N][i])
			}
		}
		return temp
	}

	ret := &dp[startColor][house][currentColor]
	if *ret != 0 {
		return *ret
	}

	*ret = INF

	for i := 1; i <= 3; i++ {
		// i+1번째 집을 칠하는 색은 i번째 집을 칠한 색과 달라야 한다
		if currentColor != i {
			*ret = min(*ret, solve(house+1, i, startColor)+cost[house+1][i])
		}
	}

	return *ret
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
