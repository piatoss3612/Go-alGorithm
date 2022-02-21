package bj2133

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

	dp := make([]int, 31)
	dp[0] = 1
	dp[2] = 3

	for i := 4; i <= n; i += 2 {
		dp[i] = dp[i-2] * dp[2]
		for j := 0; j <= i-4; j += 2 {
			dp[i] += dp[j] * 2
		}
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
