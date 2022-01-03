package bj15991

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
	dp := make([]int, 100001)
	// 6번째 항까지는 따로 점화식이 없음
	dp[1] = 1
	dp[2] = 2
	dp[3] = 2
	dp[4] = 3
	dp[5] = 3
	dp[6] = 6

	// dp[i-2]: 1 + 대칭 + 1인 경우
	// dp[i-4]: 2 + 대칭 + 2인 경우
	// dp[i-6]: 3 + 대칭 + 3인 경우
	for i := 7; i <= 100000; i++ {
		dp[i] = (dp[i-2] + dp[i-4] + dp[i-6]) % 1000000009
	}
	t := scanInt()
	for i := 0; i < t; i++ {
		n := scanInt()
		fmt.Fprintln(writer, dp[n])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
