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
	T, k, N int
	dp      [21][211]int // dp[k][N]: 다항식 p(x)에 주어진 k값에 대하여 x의 차수가 N인 항의 계수를 메모이제이션
)

func init() {
	// k가 1일 때 p(x) = x - 1
	dp[1][0] = 1
	dp[1][1] = 1
	for i := 2; i <= 20; i++ {
		end := i * (i + 1) / 2 // i (2 <= i <= k)가 주어졌을 때 p(x)의 최고차항의 차수

		// p(x)의 정수부와 최고차항의 계수는 항상 1
		dp[i][0] = 1
		dp[i][end] = 1

		/*
			특정 k값에 대한 p(x)의 각 항의 계수는
			k-1이 주어졌을 때 p(x)의 각 항의 계수 = dp[k-1]과 (1 + x + ... + x^k-1 + x^k)를 사용하여 구할 수 있다

			머리로는 이해하겠는데 설명을 못하겠네...
		*/

		for j := 1; j <= end/2; j++ {
			if j <= i {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-i-1]
				// 앞에서 누적된 dp[i-1][j-i-1]는 dp[i-1]의 차수가 j-i-1인 항을 사용하여 만들 수 없는 값이므로 그만큼 차감한다
			}
			dp[i][end-j] = dp[i][j] // 다항식 p(x)의 항의 계수는 대칭 형태을 이룬다
		}
	}
}

// 메모리: 940KB
// 시간: 4ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		k, N = scanInt(), scanInt()
		fmt.Fprintln(writer, dp[k][N])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
