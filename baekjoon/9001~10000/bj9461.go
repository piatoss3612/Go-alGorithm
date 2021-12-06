package bj9461

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

	dp := make([]int, 101)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= 100; i++ {
		dp[i] = dp[i-2] + dp[i-3]
	}

	n := scanInt()
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, dp[scanInt()])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
