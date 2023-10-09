package bj10211

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
	for i := 0; i < t; i++ {
		n := scanInt()
		dp := make([]int, n+1)
		max := -1000
		for i := 1; i <= n; i++ {
			dp[i] = scanInt()
			if dp[i-1] > 0 {
				dp[i] += dp[i-1]
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
		fmt.Fprintln(writer, dp)
		fmt.Fprintln(writer, max)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
