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
	for {
		n, m := scanInt(), scanInt()
		if n == 0 && m == 0 {
			return
		}

		dp := make([][]int, n+1)
		dp[0] = make([]int, m+1)
		for i := 1; i <= n; i++ {
			dp[i] = make([]int, m+1)
			for j := 1; j <= m; j++ {
				dp[i][j] = scanInt()
			}
		}

		max := 0

		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				if dp[i][j] == 1 {
					if dp[i-1][j-1] > 0 && dp[i-1][j] > 0 && dp[i][j-1] > 0 {
						dp[i][j] = getMin(getMin(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
					}
				}
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
		fmt.Fprintln(writer, max)
	}
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
