package main

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
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = scanInt()
	}

	max := dp[0]

	for i := 1; i < n; i++ {
		if dp[i-1] > 0 && dp[i]+dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}

		if max < dp[i] {
			max = dp[i]
		}
	}
	fmt.Fprintln(writer, max)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
