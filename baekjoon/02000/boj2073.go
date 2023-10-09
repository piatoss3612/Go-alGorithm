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
	D, P    int
	dp      [100001]int
	dp2     [351][100001]int
)

// 메모리: 2464KB
// 시간: 72ms
// 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	D, P = scanInt(), scanInt()

	for i := 1; i <= P; i++ {
		L, C := scanInt(), scanInt() // 파이프 길이와 용량 입력

		// j는 길이가 D인 경우부터 역순으로 탐색
		// 0부터 시작하면 같은 값을 여러 번 누적해서 더하게 된다
		for j := D; j >= 0; j-- {
			// j가 0인 경우
			if j == 0 && L <= D {
				dp[L] = max(dp[L], C)
				continue
			}
			// j가 0이 아닌 경우
			if j-L >= 0 && dp[j-L] > 0 {
				dp[j] = max(dp[j], min(dp[j-L], C))
			}
		}
	}

	fmt.Fprintln(writer, dp[D])
}

// 메모리: 548580KB
// 시간: 344ms
// 2차원 배열을 사용하여 메모리와 시간이 main 함수와 크게 차이난다
func main2() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	D, P = scanInt(), scanInt()

	for i := 1; i <= P; i++ {
		L, C := scanInt(), scanInt()
		for j := 0; j <= D; j++ {
			dp2[i][j] = max(dp2[i][j], dp2[i-1][j])
			if j == 0 && L <= D {
				dp2[i][L] = max(dp2[i][L], C)
				continue
			}
			if dp2[i-1][j] > 0 && j+L <= D {
				dp2[i][j+L] = max(dp2[i][j+L], min(dp2[i-1][j], C))
			}
		}
	}

	fmt.Fprintln(writer, dp2[P][D])
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
