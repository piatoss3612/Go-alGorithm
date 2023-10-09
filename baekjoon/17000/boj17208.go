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
	N, M, K int
	dp      [101][301][301]int
)

// 메모리: 72412KB
// 시간: 88ms
// 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M, K = scanInt(), scanInt(), scanInt()

	ans := 0

	// 3중 for문 이지만 배열의 크기가 101*301*301이므로 오래 걸리지는 않음
	for i := 1; i <= N; i++ {
		x, y := scanInt(), scanInt()
		for j := M; j >= 0; j-- {
			for k := K; k >= 0; k-- {
				// 남아있는 햄버거 j와 감자튀김 k에 대해서 i번째 주문을 선택할 수 있는 경우
				if j-x >= 0 && k-y >= 0 {
					// dp[i-1][j][k]: i번째 주문을 받기 이전 상태에서 햄버거 j와 감자튀김 k가 남아있을 때의 주문 개수
					// dp[i-1][j-x][k-y]+1: i번째 주문을 받기 이전 상태에서 i번째 주문을 받았을 때의 주문 개수
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-x][k-y]+1)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
				ans = max(ans, dp[i][j][k])
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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
