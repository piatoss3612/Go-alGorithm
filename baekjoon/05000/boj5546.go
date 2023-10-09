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
	N, K       int
	determined [101]int       // 특정 일에 먹을 파스타가 정해져 있는 경우
	dp         [101][4][3]int // [N일 동안][1,2,3번 파스타 중][최대 연속해서 2번까지 먹는 경우], 연속해서 3번 이상은 먹을 수 없다
)

// 메모리: 912KB
// 시간: 8ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()

	// 1. 특정 일에 먹을 파스타가 정해져 있는 경우 입력
	for i := 1; i <= K; i++ {
		determined[scanInt()] = scanInt()
	}

	// 2. 첫째 날 파스타를 먹는 경우의 수 초기화
	if determined[1] > 0 {

		// 2-1. 첫째 날 먹을 파스타의 종류가 정해져 있는 경우
		dp[1][determined[1]][1] = 1
	} else {

		// 2-2. 첫째 날 먹을 파스타의 종류가 정해져 있지 않은 경우
		dp[1][1][1] = 1
		dp[1][2][1] = 1
		dp[1][3][1] = 1
	}

	// 3. 둘째 날부터 N번째 날까지 어떤 파스타를 먹을 수 있는지 경우의 수 탐색
	for i := 2; i <= N; i++ {

		if kind := determined[i]; kind > 0 {

			// 3-1. i번째 날에 먹을 파스타의 종류가 정해져 있는 경우
			for j := 1; j <= 3; j++ {
				if j != kind {

					// 3-1-1. i-1번째 날에 kind가 아닌 파스타를 먹는 경우
					// kind와 다른 종류의 파스타 이므로 연속해서 1번 먹은 경우와 연속해서 2번 먹은 경우를 모두 카운팅
					dp[i][kind][1] += dp[i-1][j][1] + dp[i-1][j][2]
					dp[i][kind][1] %= 10000 // 값이 커질 수 있으므로 mod 연산을 실행
				} else {

					// 3-1-2. i번째 날과 i-1번째 날에 동일한 kind 파스타를 먹는 경우
					// 연속해서 3번 이상 같은 파스타를 먹을 수 없으므로 연속해서 1번 먹을 경우만 카운팅
					dp[i][kind][2] += dp[i-1][kind][1]
				}
			}
		} else {

			// 3-2. i번째 날에 먹을 파스타의 종류가 정해져 있지 않은 경우
			for j := 1; j <= 3; j++ {
				for k := 1; k <= 3; k++ {
					if j != k {

						// 3-2-1. i번째 날과 i-1번째 날에 다른 종류의 파스타를 먹는 경우
						// j와 k는 다른 종류의 파스타 이므로 연속해서 1번 먹은 경우와 연속해서 2번 먹은 경우를 모두 카운팅
						dp[i][j][1] += dp[i-1][k][1] + dp[i-1][k][2]
						dp[i][j][1] %= 10000 // 값이 커질 수 있으므로 mod 연산을 실행
					} else {

						// 3-2-2. i번째 날과 i-1번째 날에 동일한 종류의 파스타를 먹는 경우
						// 연속해서 3번 이상 같은 파스타를 먹을 수 없으므로 연속해서 1번 먹을 경우만 카운팅
						dp[i][j][2] += dp[i-1][j][1]
					}
				}
			}
		}
	}

	ans := 0

	// 4. N일 동안 파스타를 먹는 계획의 수 구하기
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			ans += dp[N][i][j]
		}
	}

	fmt.Fprintln(writer, ans%10000)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
