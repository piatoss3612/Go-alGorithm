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
	land    [1001][1001]int // 목장을 지을 땅
	dp      [1001][1001]int // dp[i][j]: i,j 위치에서 목장의 한 변의 최대 크기
)

// 메모리: 20448KB
// 시간: 68ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			land[i][j] = scanInt()
		}
	}

	ans := 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			// 장애물이 없는 경우
			if land[i][j] == 0 {
				// (i,j) 위치에서의 최댓값은 각각의 좌표 (i-1, j-1), (i-1, j), (i, j-1)에
				// 메모이제이션된 정사각형의 한 변의 크기의 최댓값들의 최솟값에 1을 더해준 것
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				ans = max(ans, dp[i][j])
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
