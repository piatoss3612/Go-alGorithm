package bj9095

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
	t := scanInt()
	dp := make([]int, 12)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4

	for i := 4; i <= 11; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

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
