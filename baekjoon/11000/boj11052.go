package bj11052

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
	cost := make([]int, n+1)
	for i := 1; i <= n; i++ {
		cost[i] = scanInt()
	}
	dp := make([]int, n+1)
	dp[1] = cost[1] // dp[1]의 최댓값은 cost[1]
	for i := 2; i <= n; i++ {
		dp[i] = cost[i]          // 임시 최댓값
		for j := 1; j < i; j++ { // 1부터 i - 1까지 반복하면서 최댓값 찾기
			tmp := dp[j] + dp[i-j]
			if tmp > dp[i] {
				dp[i] = tmp
			}
		}
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
