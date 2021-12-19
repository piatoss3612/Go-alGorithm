package bj9507

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
	dp := make([]int, 68)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4

	for i := 4; i <= 67; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3] + dp[i-4]
	}

	t := scanInt()
	for i := 0; i < t; i++ {
		fmt.Fprintln(writer, dp[scanInt()])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
