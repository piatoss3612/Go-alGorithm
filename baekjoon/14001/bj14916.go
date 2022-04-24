package bj14916

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

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if i%5 == 0 {
			dp[i] = dp[i-5] + 1
		} else if i%2 == 0 {
			dp[i] = dp[i-2] + 1
		} else if i > 5 {
			if dp[i-5] > dp[i-2] {
				dp[i] = dp[i-2] + 1
			} else {
				dp[i] = dp[i-5] + 1
			}
		} else {
			dp[i] = -1
		}
	}
	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
