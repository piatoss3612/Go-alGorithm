package bj2410

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	dp := make([]int, 1000001)

	dp[1] = 1 // 1
	dp[2] = 2 // 1+1, 2
	dp[3] = 2 // 1+1+1, 1+2

	/*
		인덱스가 2로 나누어 떨어지는 경우
		dp[i] = dp[i - 1] + dp[i / 2]

		예시:
		dp[4] = dp[3] + dp[2] // 1+1+1+1, 1+1+2, 2+2, 4
		1+1+1+1, 1+1+2: dp[3]의 모든 경우의 수에 1을 더한 수
		2+2, 4: dp[2]의 모든 경우의 수에 2를 곱한 수
	*/
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1]
		if i%2 == 0 {
			dp[i] = (dp[i] + dp[i/2]) % 1000000000
		}
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
