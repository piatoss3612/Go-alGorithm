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
	N, K    int
	dp      [5001]int
)

// 난이도: Gold 4
// 메모리: 196272KB -> 988KB
// 시간: 232ms -> 140ms
// 분류: 다이나믹 프로그래밍, 조합론
// 2차원 dp를 출력해본 결과, N과 K값이 변하더라도 메모이제이션된 값은 항상 동일하며 일정한 규칙이 존재하므로
// 이를 바탕으로 dp[i][j] = dp[i-1][j] + dp[i][j-1]이라는 점화식을 발견할 수 있었다
// 그리고 이 점화식을 적용하여 dp를 1차원 슬라이스로 단순화하여 시간과 메모리를 절약하였다

func main() {
	defer writer.Flush()
	Input()
	Solve()
}

const MOD = 1000000000

func Input() {
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()
}

func Solve() {
	dp[0] = 1
	for i := 1; i <= K; i++ {

		for j := 1; j <= N; j++ {
			dp[j] += dp[j-1]
			dp[j] %= MOD
		}
	}

	fmt.Fprintln(writer, dp[N])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
