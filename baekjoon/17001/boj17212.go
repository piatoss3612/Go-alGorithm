package bj17212

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
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i >= 2 {
			if dp[i] > dp[i-2]+1 {
				dp[i] = dp[i-2] + 1
			}
			if i >= 5 {
				if dp[i] > dp[i-5]+1 {
					dp[i] = dp[i-5] + 1
				}
				if i >= 7 {
					if dp[i] > dp[i-7]+1 {
						dp[i] = dp[i-7] + 1
					}
				}
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
