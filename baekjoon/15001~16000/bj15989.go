package bj15989

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
	dp := make([][4]int, 10001)
	dp[1] = [4]int{0, 1, 0, 0} // 1
	dp[2] = [4]int{0, 1, 1, 0} // 1 + 1, 2
	dp[3] = [4]int{0, 1, 1, 1} // 1 + 1 + 1, 2 + 1, 3
	for i := 4; i <= 10000; i++ {
		dp[i][1] = dp[i-1][1]                           // 1 + .. 로 시작하는 합
		dp[i][2] = dp[i-2][1] + dp[i-2][2]              // 2 + .. 로 시작하는 합
		dp[i][3] = dp[i-3][1] + dp[i-3][2] + dp[i-3][3] // 3 + .. 로 시작하는 합
	}
	t := scanInt()
	for i := 0; i < t; i++ {
		n := scanInt()
		fmt.Fprintln(writer, dp[n][1]+dp[n][2]+dp[n][3])
	}

}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
