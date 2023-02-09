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
	dp      [1000001]int
	T, N    int
)

func init() {
	// dp[i] = 1인 경우: 구사과의 승리
	// dp[i] = 0인 경우: 큐브러버의 승리 = 구사과의 패배

	dp[1] = 1 // 1에서 시작하는 경우는 항상 구사과의 승리
	for i := 2; i <= 1000000; i++ {
		win := false
		for j := 1; j*j <= i; j++ {
			// 항상 구사과가 먼저 시작하므로 구사과가 승리할 수 있는 경우의 수를 기준으로 탐색
			// 정수 i에서 게임을 시작했을 때 구사과가 이기는 경우의 수가 존재할 경우
			if dp[i-j*j] == 0 {
				win = true
				break
			}
		}

		if win {
			dp[i] = 1 // 구사과의 승리
		} else {
			dp[i] = 0 // 구사과의 패배
		}
	}
}

// 난이도: Gold 5
// 메모리: 9548KB
// 시간: 248ms
// 분류: 다이나믹 프로그래밍, 게임 이론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		N = scanInt()
		if dp[N] == 1 {
			fmt.Fprintln(writer, "koosaga")
		} else {
			fmt.Fprintln(writer, "cubelover")
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
