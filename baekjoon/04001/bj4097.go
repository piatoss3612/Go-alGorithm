package bj4097

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
	for {
		n := scanInt()
		if n == 0 {
			break
		} else {
			runTest(n)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func runTest(n int) {
	max := -10000
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = scanInt()
		if dp[i-1]+dp[i] > dp[i] {
			dp[i] += dp[i-1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	fmt.Fprintln(writer, max)
}
