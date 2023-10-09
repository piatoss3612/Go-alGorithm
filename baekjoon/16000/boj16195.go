package bj16195

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
	dp := make([][]int, 1001)
	for i := 0; i <= 1000; i++ {
		dp[i] = make([]int, 1001)
	}

	dp[1][1] = 1
	dp[2][1] = 1
	dp[2][2] = 1
	dp[3][1] = 1
	dp[3][2] = 2
	dp[3][3] = 1

	for i := 4; i <= 1000; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = (dp[i-3][j-1] + dp[i-2][j-1] + dp[i-1][j-1]) % 1000000009
		}
	}

	t := scanInt()

	for i := 0; i < t; i++ {
		n, m := scanInt(), scanInt()
		ans := 0
		for i := 1; i <= m; i++ {
			ans = (ans + dp[n][i]) % 1000000009
		}
		fmt.Fprintln(writer, ans)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
